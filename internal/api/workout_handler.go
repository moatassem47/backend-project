package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/moatassem47/back-end-project/internal/store"
	"github.com/moatassem47/back-end-project/internal/utils"
)

type WorkoutHandler struct {
	workoutStore store.WorkoutStore
	logger       *log.Logger
}

func NewWorkoutHandler(workoutStore store.WorkoutStore, logger *log.Logger) *WorkoutHandler {
	return &WorkoutHandler{
		workoutStore: workoutStore,
		logger:       logger,
	}
}
func (wh *WorkoutHandler) HandleGetWorkoutById(w http.ResponseWriter, r *http.Request) {
	workoutID, err := utils.ReadIDParam(r)
	if err != nil {
		wh.logger.Printf("ERROR: read ID param:%v", err)
		utils.WriteJson(w, http.StatusBadRequest, utils.Envelope{"error": "invalid workout ID"})
		return

	}
	workout, err := wh.workoutStore.GetWorkoutByID(workoutID)
	if err != nil {
		wh.logger.Printf("ERROR:getWorkoutByID:%v", err)
		utils.WriteJson(w, http.StatusInternalServerError, utils.Envelope{"error": "Internal Server Error"})
		return
	}
	utils.WriteJson(w, http.StatusOK, utils.Envelope{"workout": workout})
}
func (wh *WorkoutHandler) HandleCreateWorkout(w http.ResponseWriter, r *http.Request) {
	var workout store.Workout
	err := json.NewDecoder(r.Body).Decode(&workout)
	if err != nil {
		wh.logger.Printf("ERROR: decoding create workout:%v", err)
		utils.WriteJson(w, http.StatusBadRequest, utils.Envelope{"error": "invalid request sent"})
		return
	}
	createdWorkout, err := wh.workoutStore.CreateWorkout(&workout)
	if err != nil {
		wh.logger.Printf("ERROR: CreateWorkout:%v", err)
		utils.WriteJson(w, http.StatusInternalServerError, utils.Envelope{"error": "failed to create workout"})
		return
	}
	utils.WriteJson(w, http.StatusCreated, utils.Envelope{"workout": createdWorkout})
}
func (wh *WorkoutHandler) HandleUpdateWorkoutByID(w http.ResponseWriter, r *http.Request) {
	workoutID, err := utils.ReadIDParam(r)
	if err != nil {
		wh.logger.Printf("ERROR: read id param:%v", err)
		utils.WriteJson(w, http.StatusBadRequest, utils.Envelope{"error": "invalid update"})
		return
	}

	existingWorkout, err := wh.workoutStore.GetWorkoutByID(workoutID)
	if err != nil {
		wh.logger.Printf("ERROR: getWorkoutByID:%v", err)
		utils.WriteJson(w, http.StatusInternalServerError, utils.Envelope{"error": "internal server error"})
		return
	}
	if existingWorkout == nil {
		http.NotFound(w, r)
		return
	}
	var updateWorkoutRequest struct {
		Title           *string              `json:"title"`
		Description     *string              `json:"description"`
		DurationMinutes *int                 `json:"duration_minutes"`
		CaloriesBurnt   *int                 `json:"calories_burnt"`
		Entries         []store.WorkoutEntry `json:"entries"`
	}
	err = json.NewDecoder(r.Body).Decode(&updateWorkoutRequest)
	if err != nil {
		wh.logger.Printf("ERROR: decoding request:%v", err)
		utils.WriteJson(w, http.StatusBadRequest, utils.Envelope{"error": "invalid request"})

	}

	if updateWorkoutRequest.Title != nil {
		existingWorkout.Title = *updateWorkoutRequest.Title
	}

	if updateWorkoutRequest.Description != nil {
		existingWorkout.Description = *updateWorkoutRequest.Description
	}

	if updateWorkoutRequest.DurationMinutes != nil {
		existingWorkout.DurationMinutes = *updateWorkoutRequest.DurationMinutes
	}

	if updateWorkoutRequest.CaloriesBurnt != nil {
		existingWorkout.CaloriesBurnt = *updateWorkoutRequest.CaloriesBurnt
	}

	if updateWorkoutRequest.Entries != nil {
		existingWorkout.Entries = updateWorkoutRequest.Entries
	}
	err = wh.workoutStore.UpdateWorkout(existingWorkout)
	if err != nil {

		wh.logger.Printf("ERROR:UpdateWorkout:%v", err)
		utils.WriteJson(w, http.StatusInternalServerError, utils.Envelope{"error": "internal server error"})
	}
	utils.WriteJson(w, http.StatusOK, utils.Envelope{"wokrout": existingWorkout})
}
func (wh *WorkoutHandler) HandleDeleteWorkoutByID(w http.ResponseWriter, r *http.Request) {
	paramsWorkoutId := chi.URLParam(r, "id")
	if paramsWorkoutId == "" {
		http.NotFound(w, r)
		return
	}
	workoutId, err := strconv.ParseInt(paramsWorkoutId, 10, 64)
	if err != nil {
		wh.logger.Printf("ERROR: read id param:%v", err)
		utils.WriteJson(w, http.StatusBadRequest, utils.Envelope{"error": "invalid ID"})
	}
	err = wh.workoutStore.DeleteWorkout(workoutId)
	if err == sql.ErrNoRows {
		fmt.Println(err)
		http.Error(w, "workout not found", http.StatusNotFound)
		return
	}
	if err != nil {
		wh.logger.Printf("ERROR: DeleteWorkout:%v", err)
		utils.WriteJson(w, http.StatusBadRequest, utils.Envelope{"error": "invalide delete"})
	}
	w.WriteHeader(http.StatusNoContent)

}

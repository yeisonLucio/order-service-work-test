package models

import (
	"testing"
	"time"
)

func TestWorkOrder_Validate(t *testing.T) {
	type fields struct {
		PlannedDateBegin *time.Time
		PlannedDateEnd   *time.Time
		Type             string
		Title            string
	}
	beginDate := time.Now()
	endDate := beginDate.Add(time.Hour * 3)
	correctEndDate := beginDate.Add(time.Hour)

	tests := []struct {
		name          string
		fields        fields
		wantErr       bool
		expectedError error
	}{
		{
			name:          "should return an error when title is empty",
			fields:        fields{},
			wantErr:       true,
			expectedError: ErrInvalidTitle,
		},
		{
			name: "should return an error when type is not valid",
			fields: fields{
				Title: "title",
				Type:  "fake",
			},
			wantErr:       true,
			expectedError: ErrInvalidType,
		},
		{
			name: "should return an error when range date is not valid",
			fields: fields{
				Title:            "title",
				Type:             ServiceOrderType,
				PlannedDateBegin: &beginDate,
				PlannedDateEnd:   &endDate,
			},
			wantErr:       true,
			expectedError: ErrInvalidDates,
		},
		{
			name: "should return nil when validations pass",
			fields: fields{
				Title:            "title",
				Type:             ServiceOrderType,
				PlannedDateBegin: &beginDate,
				PlannedDateEnd:   &correctEndDate,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WorkOrder{
				PlannedDateBegin: tt.fields.PlannedDateBegin,
				PlannedDateEnd:   tt.fields.PlannedDateEnd,
				Type:             tt.fields.Type,
				Title:            tt.fields.Title,
			}
			err := w.Validate()

			if (err != nil) != tt.wantErr {
				t.Errorf("WorkOrder.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err != nil && err != tt.expectedError {
				t.Errorf("WorkOrder.Validate() error = %v, expectedErr %v", err, tt.expectedError)
			}
		})
	}
}

func TestWorkOrder_SetPlannedDateBegin(t *testing.T) {

	type args struct {
		date string
	}
	tests := []struct {
		name          string
		args          args
		expectedError error
		wantErr       bool
	}{
		{
			name: "should return an error when date is not valid",
			args: args{
				date: "fake",
			},
			wantErr:       true,
			expectedError: ErrBeginDateFormat,
		},
		{
			name: "should set date successfully when format is ok",
			args: args{
				date: "2023-07-03 00:00:00",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WorkOrder{}
			err := w.SetPlannedDateBegin(tt.args.date)

			if (err != nil) != tt.wantErr {
				t.Errorf("WorkOrder.SetPlannedDateBegin() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err != nil && err != tt.expectedError {
				t.Errorf("WorkOrder.Validate() error = %v, expectedErr %v", err, tt.expectedError)
			}

			if !tt.wantErr && w.PlannedDateBegin == nil {
				t.Errorf("WorkOrder.PlannedDateBegin is nil")
			}
		})
	}
}

func TestWorkOrder_SetPlannedDateEnd(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name        string
		args        args
		expectedErr error
		wantErr     bool
	}{
		{
			name: "should return an error when format is invalid",
			args: args{
				date: "invalid",
			},
			expectedErr: ErrEndDateFormat,
			wantErr:     true,
		},
		{
			name: "should set end date successfully when format is ok",
			args: args{
				date: "2023-07-03 00:00:00",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WorkOrder{}
			err := w.SetPlannedDateEnd(tt.args.date)

			if (err != nil) != tt.wantErr {
				t.Errorf("WorkOrder.SetPlannedDateEnd() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err != nil && err != tt.expectedErr {
				t.Errorf("WorkOrder.Validate() error = %v, expectedErr %v", err, tt.expectedErr)
			}

			if !tt.wantErr && w.PlannedDateEnd == nil {
				t.Errorf("WorkOrder.PlannedDateEnd is nil")
			}
		})
	}
}

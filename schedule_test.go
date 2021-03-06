package huego

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSchedules(t *testing.T) {
	b := New(hostname, username, clientKey)
	schedules, err := b.GetSchedules()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Found %d schedules", len(schedules))
	for i, schedule := range schedules {
		t.Logf("%d:", i)
		t.Logf("  Name: %s", schedule.Name)
		t.Logf("  Description: %s", schedule.Description)
		t.Logf("  Command:")
		t.Logf("    Address: %s", schedule.Command.Address)
		t.Logf("    Method: %s", schedule.Command.Method)
		t.Logf("    Body: %s", schedule.Command.Body)
		t.Logf("  Time: %s", schedule.Time)
		t.Logf("  LocalTime: %s", schedule.LocalTime)
		t.Logf("  StartTime: %s", schedule.StartTime)
		t.Logf("  Status: %s", schedule.Status)
		t.Logf("  AutoDelete: %t", schedule.AutoDelete)
		t.Logf("  ID: %d", schedule.ID)
	}

	contains := func(name string, ss []*Schedule) bool {
		for _, s := range ss {
			if s.Name == name {
				return true
			}
		}
		return false
	}

	assert.True(t, contains("Timer", schedules))
	assert.True(t, contains("Alarm", schedules))

	b.Host = badHostname
	_, err = b.GetSchedules()
	assert.NotNil(t, err)
}

func TestGetSchedule(t *testing.T) {
	b := New(hostname, username, clientKey)
	s, err := b.GetSchedule(1)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Time: %s", s.Time)
	t.Logf("LocalTime: %s", s.LocalTime)
	t.Logf("StartTime: %s", s.StartTime)
	t.Logf("Status: %s", s.Status)
	t.Logf("AutoDelete: %t", s.AutoDelete)
	t.Logf("ID: %d", s.ID)

	b.Host = badHostname
	_, err = b.GetSchedule(1)
	assert.NotNil(t, err)
}

func TestCreateSchedule(t *testing.T) {
	b := New(hostname, username, clientKey)
	command := &Command{
		Address: "/api/" + username + "/lights/0",
		Body: &struct {
			on bool
		}{
			false,
		},
		Method: "PUT",
	}
	schedule := &Schedule{
		Name:        "TestSchedule",
		Description: "Huego test schedule",
		Command:     command,
		LocalTime:   "2019-09-22T13:37:00",
	}
	resp, err := b.CreateSchedule(schedule)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("Schedule created")
		for k, v := range resp.Success {
			t.Logf("%v: %s", k, v)
		}
	}

	b.Host = badHostname
	_, err = b.CreateSchedule(schedule)
	assert.NotNil(t, err)
}

func TestUpdateSchedule(t *testing.T) {
	b := New(hostname, username, clientKey)
	id := 1
	schedule := &Schedule{
		Name:        "New Scehdule",
		Description: "Updated parameter",
	}
	resp, err := b.UpdateSchedule(id, schedule)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("Schedule %d updated", id)
		for k, v := range resp.Success {
			t.Logf("%v: %s", k, v)
		}
	}

	b.Host = badHostname
	_, err = b.UpdateSchedule(id, schedule)
	assert.NotNil(t, err)
}

func TestDeleteSchedule(t *testing.T) {
	b := New(hostname, username, clientKey)
	id := 1
	err := b.DeleteSchedule(id)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("Schedule %d deleted", id)
	}
}

package main

import "fmt"

// Question
// Given a patient(diagnoses and medical history), and a drug that has requirement to be able to use it;
// create a function that tells if that patient qualifies for that drug or not.

// Example
// Drug A required the patient to have diabetes, and to have tried "old drug" for at least 42 days
// Patient A has diabeted and has been on "old drug" for 50 days
// Answer: True

// Ike's Synopsis afterwards
// Most likley failed. Initially I thought I got to a working solution, but based on memory I don't think
// I checked the condition when assessing requriemnts. I only checked the time on the previous drug.
// Generally I just think I over-complicated it, but I don't know how to siplify it in such a small amount of time.
// Only had about 30 minutes (interview started at 1, chatted for 5 min, setup env, stopped working at 1:40 and then
// chatted with her for 5 minutes). Maybe next time I should start with the function, and sudo-code what I need it to do,
// before I go into modeling the types. Modeling the types is something that's a bit difficult to do before you know enough.

type Medication struct {
	Name                       string
	RequiredPreviousDrug       string
	RequiredTimeOnPreviousDrug int
	RequiredDiagnosis          string
}

var Medications = map[string]Medication{}

type Patient struct {
	Diagnoses         map[string]struct{}
	MedicationHistory map[string]int // This is a map from medication to the amount of time the patient was on it
}

func (p *Patient) CheckIfAllowed(drugName string) bool {
	// Get medication's information
	med, ok := Medications[drugName]
	if !ok {
		return false
	}

	// Check if patient has medication's required diagnosis
	_, ok = p.Diagnoses[med.RequiredDiagnosis]
	if !ok {
		return false
	}

	// Check if patient has taken prior medication
	daysOnDrug, ok := p.MedicationHistory[med.RequiredPreviousDrug]
	if !ok {
		return false
	}

	// Check if patient has taken prior medication for enough time
	if daysOnDrug < med.RequiredTimeOnPreviousDrug {
		return false
	}

	return true
}

func main() {
	medicationA := Medication{ //Should succeed
		Name:                       "med_a",
		RequiredPreviousDrug:       "med_old",
		RequiredTimeOnPreviousDrug: 10,
		RequiredDiagnosis:          "sick",
	}
	medicationB := Medication{ //Should fail on required diagnosis
		Name:                       "med_b",
		RequiredPreviousDrug:       "med_old",
		RequiredTimeOnPreviousDrug: 10,
		RequiredDiagnosis:          "fake_diagnosis", // fails here
	}
	medicationC := Medication{ //Should fail on required time on previous drug
		Name:                       "med_c",
		RequiredPreviousDrug:       "med_old",
		RequiredTimeOnPreviousDrug: 100, // fails here
		RequiredDiagnosis:          "injury",
	}
	medicationD := Medication{ //Should fail on previous drug
		Name:                       "med_d",
		RequiredPreviousDrug:       "fake_med", // fails here
		RequiredTimeOnPreviousDrug: 10,
		RequiredDiagnosis:          "injury",
	}

	Medications[medicationA.Name] = medicationA
	Medications[medicationB.Name] = medicationB
	Medications[medicationC.Name] = medicationC
	Medications[medicationD.Name] = medicationD

	diagnoses := map[string]struct{}{
		"injury": {},
		"sick":   {},
	}

	history := map[string]int{
		"med_old": 45,
	}

	patient := Patient{
		Diagnoses:         diagnoses,
		MedicationHistory: history,
	}

	res1 := patient.CheckIfAllowed("med_a")
	res2 := patient.CheckIfAllowed("med_b")
	res3 := patient.CheckIfAllowed("med_c")
	res4 := patient.CheckIfAllowed("med_d")

	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)
	fmt.Println(res4)
}

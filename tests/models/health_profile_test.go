package models_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/adil-faiyaz98/go-builder-kit/builders"
	"github.com/adil-faiyaz98/go-builder-kit/models"
)

var _ = Describe("HealthProfile", func() {
	var (
		healthProfileBuilder *builders.HealthProfileBuilder
		insuranceBuilder     *builders.InsuranceBuilder
		medicationBuilder    *builders.MedicationBuilder
		personBuilder        *builders.PersonBuilder
		addressBuilder       *builders.AddressBuilder
		medicalRecordBuilder *builders.MedicalRecordBuilder
	)

	BeforeEach(func() {
		addressBuilder = builders.NewAddressBuilder().
			WithStreet("123 Medical Center Dr").
			WithCity("Boston").
			WithState("MA").
			WithPostalCode("02115").
			WithCountry("USA").
			WithType("Business")

		personBuilder = builders.NewPersonBuilder().
			WithID("doctor-123").
			WithName("Dr. Jane Smith").
			WithEmail("dr.smith@hospital.com").
			WithPhone("+1-555-123-4567")

		medicationBuilder = builders.NewMedicationBuilder().
			WithName("Lisinopril").
			WithDosage("10mg").
			WithFrequency("Once daily").
			WithStartDate("2022-01-15").
			WithPrescribedBy("Dr. Jane Smith")

		medicalRecordBuilder = builders.NewMedicalRecordBuilder().
			WithDate("2022-01-15").
			WithType("Checkup").
			WithProvider("Boston Medical Center").
			WithDiagnosis("Hypertension").
			WithTreatment("Prescribed Lisinopril").
			WithLocation(addressBuilder)

		insuranceBuilder = builders.NewInsuranceBuilder().
			WithProvider("Health Insurance Co").
			WithPolicyNumber("POL-123456").
			WithType("Health").
			WithStartDate("2023-01-01").
			WithEndDate("2023-12-31")

		healthProfileBuilder = builders.NewHealthProfileBuilder().
			WithHeight(180.0).
			WithWeight(75.0).
			WithBloodType("O+").
			WithAllergy("Peanuts").
			WithAllergy("Penicillin").
			WithChronicCondition("Hypertension").
			WithMedication(medicationBuilder).
			WithMedicalRecord(medicalRecordBuilder).
			WithPrimaryPhysician(personBuilder).
			WithInsurance(insuranceBuilder)
	})

	Describe("Positive Tests", func() {
		Context("when creating a health profile with valid data", func() {
			It("should create a health profile with basic information", func() {
				healthProfile := healthProfileBuilder.Build().(*models.HealthProfile)

				Expect(healthProfile.Height).To(Equal(180.0))
				Expect(healthProfile.Weight).To(Equal(75.0))
				Expect(healthProfile.BloodType).To(Equal("O+"))
				Expect(len(healthProfile.Allergies)).To(Equal(2))
				Expect(healthProfile.Allergies).To(ContainElements("Peanuts", "Penicillin"))
			})

			It("should create a health profile with chronic conditions", func() {
				healthProfile := healthProfileBuilder.
					WithChronicCondition("Asthma").
					Build().(models.HealthProfile)

				Expect(len(healthProfile.ChronicConditions)).To(Equal(2))
				Expect(healthProfile.ChronicConditions).To(ContainElements("Hypertension", "Asthma"))
			})

			It("should create a health profile with medications", func() {
				healthProfile := healthProfileBuilder.Build().(*models.HealthProfile)

				Expect(len(healthProfile.Medications)).To(Equal(1))
				Expect(healthProfile.Medications[0].Name).To(Equal("Lisinopril"))
				Expect(healthProfile.Medications[0].Dosage).To(Equal("10mg"))
				Expect(healthProfile.Medications[0].Frequency).To(Equal("Once daily"))
			})

			It("should create a health profile with primary physician", func() {
				healthProfile := healthProfileBuilder.Build().(*models.HealthProfile)

				// Since PrimaryPhysician is of type any, we need to cast it
				physician := healthProfile.PrimaryPhysician.(*models.Person)
				Expect(physician.Name).To(Equal("Dr. Jane Smith"))
				Expect(physician.Email).To(Equal("dr.smith@hospital.com"))
			})

			It("should create a health profile with insurance", func() {
				healthProfile := healthProfileBuilder.Build().(*models.HealthProfile)

				Expect(healthProfile.Insurance.Provider).To(Equal("Health Insurance Co"))
				Expect(healthProfile.Insurance.PolicyNumber).To(Equal("POL-123456"))
				Expect(healthProfile.Insurance.Type).To(Equal("Health"))
			})

			It("should calculate BMI correctly", func() {
				// BMI = weight(kg) / (height(m) * height(m))
				// For 75kg and 180cm (1.8m): 75 / (1.8 * 1.8) = 23.15
				healthProfile := healthProfileBuilder.Build().(*models.HealthProfile)

				// If we had a BMI calculation method, we would test it here
				// For now, we'll just calculate it manually
				bmi := healthProfile.Weight / ((healthProfile.Height / 100) * (healthProfile.Height / 100))
				Expect(bmi).To(BeNumerically("~", 23.15, 0.01))
			})
		})
	})

	Describe("Negative Tests", func() {
		Context("when creating a health profile with invalid data", func() {
			It("should handle negative height", func() {
				healthProfile := healthProfileBuilder.
					WithHeight(-180.0).
					Build().(models.HealthProfile)

				// If we had validation, this would fail
				// For now, just check the height is negative
				Expect(healthProfile.Height).To(Equal(-180.0))
			})

			It("should handle negative weight", func() {
				healthProfile := healthProfileBuilder.
					WithWeight(-75.0).
					Build().(models.HealthProfile)

				// If we had validation, this would fail
				// For now, just check the weight is negative
				Expect(healthProfile.Weight).To(Equal(-75.0))
			})

			It("should handle invalid blood type", func() {
				healthProfile := healthProfileBuilder.
					WithBloodType("XYZ"). // Invalid blood type
					Build().(models.HealthProfile)

				// If we had validation, this would fail
				// For now, just check the blood type
				Expect(healthProfile.BloodType).To(Equal("XYZ"))
			})

			It("should handle invalid medication data", func() {
				invalidMedicationBuilder := builders.NewMedicationBuilder().
					WithName("").     // Invalid: empty name
					WithDosage("").   // Invalid: empty dosage
					WithFrequency("") // Invalid: empty frequency

				healthProfile := healthProfileBuilder.
					WithMedication(invalidMedicationBuilder).
					Build().(models.HealthProfile)

				// If we had validation, this would fail
				// For now, just check the medication data
				Expect(healthProfile.Medications[1].Name).To(Equal(""))
				Expect(healthProfile.Medications[1].Dosage).To(Equal(""))
				Expect(healthProfile.Medications[1].Frequency).To(Equal(""))
			})

			It("should handle invalid insurance data", func() {
				invalidInsuranceBuilder := builders.NewInsuranceBuilder().
					WithProvider("").     // Invalid: empty provider
					WithPolicyNumber(""). // Invalid: empty policy number
					WithStartDate("2023-01-01").
					WithEndDate("2022-01-01") // Invalid: end date before start date

				healthProfile := healthProfileBuilder.
					WithInsurance(invalidInsuranceBuilder).
					Build().(models.HealthProfile)

				// If we had validation, this would fail
				// For now, just check the insurance data
				Expect(healthProfile.Insurance.Provider).To(Equal(""))
				Expect(healthProfile.Insurance.PolicyNumber).To(Equal(""))
				Expect(healthProfile.Insurance.StartDate).To(Equal("2023-01-01"))
				Expect(healthProfile.Insurance.EndDate).To(Equal("2022-01-01"))
			})
		})
	})

	Describe("Edge Cases", func() {
		It("should handle a health profile with no allergies", func() {
			healthProfile := builders.NewHealthProfileBuilder().
				WithHeight(175.0).
				WithWeight(70.0).
				WithBloodType("A+").
				Build().(models.HealthProfile)

			Expect(len(healthProfile.Allergies)).To(Equal(0))
			Expect(healthProfile.Allergies).To(BeEmpty())
		})

		It("should handle a health profile with many medications", func() {
			manyMedsBuilder := builders.NewHealthProfileBuilder().
				WithHeight(175.0).
				WithWeight(70.0).
				WithBloodType("A+")

			// Add 5 medications
			for i := 1; i <= 5; i++ {
				medicationBuilder := builders.NewMedicationBuilder().
					WithName(fmt.Sprintf("Medication %d", i)).
					WithDosage(fmt.Sprintf("%dmg", i*10)).
					WithFrequency("Daily")

				manyMedsBuilder.WithMedication(medicationBuilder)
			}

			healthProfile := manyMedsBuilder.Build().(*models.HealthProfile)

			Expect(len(healthProfile.Medications)).To(Equal(5))
			Expect(healthProfile.Medications[0].Name).To(Equal("Medication 1"))
			Expect(healthProfile.Medications[0].Dosage).To(Equal("10mg"))
			Expect(healthProfile.Medications[4].Name).To(Equal("Medication 5"))
			Expect(healthProfile.Medications[4].Dosage).To(Equal("50mg"))
		})

		It("should handle a health profile with extreme height and weight", func() {
			// World record values
			extremeHealthProfile := builders.NewHealthProfileBuilder().
				WithHeight(272.0). // Robert Wadlow, tallest person (272 cm)
				WithWeight(635.0). // Jon Brower Minnoch, heaviest person (635 kg)
				WithBloodType("AB-").
				Build().(models.HealthProfile)

			Expect(extremeHealthProfile.Height).To(Equal(272.0))
			Expect(extremeHealthProfile.Weight).To(Equal(635.0))

			// BMI would be extremely high
			bmi := extremeHealthProfile.Weight / ((extremeHealthProfile.Height / 100) * (extremeHealthProfile.Height / 100))
			Expect(bmi).To(BeNumerically(">", 50))
		})

		It("should handle a health profile with multiple medical records", func() {
			record1 := builders.NewMedicalRecordBuilder().
				WithDate("2022-01-15").
				WithType("Checkup").
				WithProvider("Dr. Smith")

			record2 := builders.NewMedicalRecordBuilder().
				WithDate("2022-02-15").
				WithType("Follow-up").
				WithProvider("Dr. Johnson")

			healthProfile := builders.NewHealthProfileBuilder().
				WithMedicalRecord(record1).
				WithMedicalRecord(record2).
				Build().(models.HealthProfile)

			Expect(len(healthProfile.MedicalHistory)).To(Equal(2))
			Expect(healthProfile.MedicalHistory[0].Provider).To(Equal("Dr. Smith"))
			Expect(healthProfile.MedicalHistory[1].Provider).To(Equal("Dr. Johnson"))
			Expect(healthProfile.MedicalHistory[0].Type).To(Equal("Checkup"))
			Expect(healthProfile.MedicalHistory[1].Type).To(Equal("Follow-up"))
		})
	})
})


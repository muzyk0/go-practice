package relationship

import (
	"testing"
)

type NewPercon struct {
	r Relationship
	p Person
}

func TestAddNew(t *testing.T) {
	tests := []struct {
		name           string
		existedMembers map[Relationship]Person
		newPerson      NewPercon
		ok             bool
	}{
		{
			name: "add mother",
			existedMembers: map[Relationship]Person{
				Father: {
					FirstName: "Misha",
					LastName:  "Popov",
					Age:       56,
				},
			},
			newPerson: NewPercon{
				r: Mother,
				p: Person{
					FirstName: "Nastya",
					LastName:  "Popova",
					Age:       50,
				},
			},
			ok: true,
		},
		{
			name: "catch replace father",
			existedMembers: map[Relationship]Person{
				Father: {
					FirstName: "Misha",
					LastName:  "Popov",
					Age:       56,
				},
			},
			newPerson: NewPercon{
				r: Father,
				p: Person{
					FirstName: "Drug",
					LastName:  "Mishi",
					Age:       57,
				},
			},
			ok: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f := &Family{
				Members: test.existedMembers,
			}
			err := f.AddNew(test.newPerson.r, test.newPerson.p)
			if (err != nil) != test.ok {
				t.Errorf("AddNew() error = %v, ok %v", err, test.ok)
			}
		})
	}
}

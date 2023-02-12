package mappers

import (
	"strconv"
	"time"

	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/entities"
)

type CurriculumMapper interface {
	CurriculumCreateMapper(schema dtos.CurriculumCreateDTO) entities.CurriculumEntity
	CurriculumCourseMapper(dtos.Curriculum, uint, uint) entities.CourseCurriculumEntity
	CurriculumFetchMapper(curriculum []dtos.CurriculumQueryReturnSchema) dtos.CurriculumFetchDTO
}

type curriculumMapper struct {
}

func NewCurriculumMapper() CurriculumMapper {
	return &curriculumMapper{}
}

func (s *curriculumMapper) CurriculumCreateMapper(curriculum dtos.CurriculumCreateDTO) entities.CurriculumEntity {
	return entities.CurriculumEntity{
		DepartmentID: curriculum.DepartmentID,
		BaseEntity: entities.BaseEntity{
			CreatedAt: time.Now().UTC(),
			IsActive:  true,
		},
	}
}

func (s *curriculumMapper) CurriculumCourseMapper(curriculum dtos.Curriculum, courseID uint, curriculumID uint) entities.CourseCurriculumEntity {
	return entities.CourseCurriculumEntity{
		CurriculumID: curriculumID,
		CourseID:     courseID,
		Semester:     curriculum.Semester,
		Year:         curriculum.Year,
		CourseLoad:   curriculum.CourseLoad,
		BaseEntity: entities.BaseEntity{
			CreatedAt: time.Now().UTC(),
			IsActive:  true,
		},
	}
}

func (s *curriculumMapper) CurriculumFetchMapper(curriculum []dtos.CurriculumQueryReturnSchema) dtos.CurriculumFetchDTO {
	curriculumInfoMap := make(map[string]dtos.CurriculumInfo)

	/**================================================================================================
	 * *                                           INFO
	 *
	 *
	 *   The make function is a built-in function in Go that is used to allocate and initialize a new
		instance of a built-in data type. In the case of maps, it is used to allocate a new, empty map and
		return a reference to it. The line curriculumInfoMap := make(map[string]CurriculumInfo) creates a
		new map with a key of type string and a value of type CurriculumInfo. The make function returns a
		reference to the newly created map, which is stored in the variable curriculumInfoMap.

		The make function is necessary to create a map because maps are reference types in Go,
		unlike slices and arrays which are value types. This means that maps must be created with the
		make function, or a new map literal, before they can be used. The make function takes two arguments:
		the type of the map (in this case map[string]CurriculumInfo), and an optional argument for the initial
		capacity of the map. If the initial capacity is not provided, it defaults to 0.

		In summary, the line curriculumInfoMap := make(map[string]CurriculumInfo) creates a new, empty map
		with a string key and a value of type CurriculumInfo, and returns a reference to it which is stored in
		the variable curriculumInfoMap.
	 *
	 *================================================================================================**/

	for _, course := range curriculum {
		key := course.Semester + "-" + strconv.Itoa(course.Year)
		curriculumInfo, exists := curriculumInfoMap[key]
		if !exists {
			curriculumInfo = dtos.CurriculumInfo{
				Semester:   course.Semester,
				Year:       course.Year,
				CourseLoad: course.CourseLoad,
				Courses:    []dtos.CourseInfo{},
			}
		}

		/**================================================================================================
		 * *                                           INFO
		 *
		 *   The if !exists block is checking whether a value for the key key already exists in the
			map curriculumInfoMap. If the key does not exist, then ok will be false and this block of
			code will be executed.

			The purpose of this block is to create a new dtos.CurriculumInfo value and add it to the map
			curriculumInfoMap under the key key. This is because we want to store the information for each
			unique combination of Semester and Year in a separate CurriculumInfo struct. By using the if !ok block,
			we ensure that a new CurriculumInfo struct is only created if the current combination of Semester and
			Year is not already in the map.


			NOTE: This code is also checking if the object for semester and year exists in the map or not.
			If it does not exist then it will create a new object and add it to the map, if it exists then it will
			just append the course to the existing object.

		 *
		 *
		 *================================================================================================**/

		curriculumInfo.Courses = append(curriculumInfo.Courses, dtos.CourseInfo{
			ID:        course.CourseID,
			Name:      course.Name,
			Code:      course.Code,
			Credits:   course.Credits,
			Ects:      course.Ects,
			CreatedAt: course.CreatedAt,
			UpdatedAt: course.UpdatedAt,
			DeletedAt: course.DeletedAt,
		})
		curriculumInfoMap[key] = curriculumInfo
	}

	var curriculumInfoArray []dtos.CurriculumInfo
	for _, value := range curriculumInfoMap {
		curriculumInfoArray = append(curriculumInfoArray, value)
	}

	return dtos.CurriculumFetchDTO{
		DepartmentID:   curriculum[0].DepartmentID,
		DepartmentCode: curriculum[0].DepartmentCode,
		DepartmentName: curriculum[0].DepartmentName,
		NumberOfYears:  curriculum[0].NumberOfYears,
		Curriculum:     curriculumInfoArray,
	}
}

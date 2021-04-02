def lowest_grade(students):
    """ print students """
    students.sort(key=lambda x: x[1])

    lowest_students = []
    lowest_value = students[0][1]

    second_lower_grade = None
    for student in students:
        if student[1] > lowest_value:
            if second_lower_grade and student[1] > second_lower_grade:
                break
            lowest_students.append(student[0])
            second_lower_grade = student[1]
    lowest_students.sort()
    return lowest_students


students = []
for _ in range(int(raw_input())):
    name = raw_input()
    score = float(raw_input())
    students.append([name, score])

for student in  lowest_grade(students):
    print student


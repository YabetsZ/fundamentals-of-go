## Task: Student Grade Calculator

Create a Go console application that allows students to calculate their average grade based on different subjects. 

### Application Features:
- Prompt the student to enter their name and the number of subjects they have taken.
- For each subject, the student should enter:
  - The subject name.
  - The grade obtained (numeric value).
- After entering all subjects and grades, the application should display:
  - The student's name.
  - Individual subject grades.
  - The calculated average grade.

### Requirements:
1. **Variables and Data Types**: Use appropriate variables and data types to store student data.
2. **Input Validation**: Use conditional statements to ensure grade values are within a valid range.
3. **Loops**: Implement loops to handle multiple subjects and grades.
4. **Collections**: Utilize collections (e.g., slices or maps) to store subject names and corresponding grades.
5. **Methods**: Define a method to calculate the average grade based on the entered grades.
6. **String Interpolation**: Use string interpolation to display the results in a user-friendly format.
7. **Testing** *(Optional)*: Write tests for your code to ensure correctness.

### Example Output:
```
Enter your name: Yabets
Enter the number of subjects you are taking: 3

--- Subject 1 ---
Enter the name of the subject: English
Enter the score in this, English, subject: 78.5

--- Subject 2 ---
Enter the name of the subject: Math
Enter the score in this, Math, subject: 99.357

--- Subject 3 ---
Enter the name of the subject: Rust    
Enter the score in this, Rust, subject: 85

--- Student Report ---
Student name: Yabets
Individual grades:
                Math: 99.36
                Rust: 85.00
             English: 78.50
    ✔️ Average Grade: 87.62
```


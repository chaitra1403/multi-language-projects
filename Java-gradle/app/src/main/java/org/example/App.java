package com.example;

import com.example.model.Student;
import com.example.service.StudentService;

import java.util.Scanner;

public class App {
    public static void main(String[] args) {

        Scanner scanner = new Scanner(System.in);
        StudentService service = new StudentService();

        while (true) {
            System.out.println("\n1. Add Student");
            System.out.println("2. View Students");
            System.out.println("3. Find Student");
            System.out.println("4. Exit");

            int choice = scanner.nextInt();

            switch (choice) {
                case 1:
                    System.out.print("ID: ");
                    int id = scanner.nextInt();
                    scanner.nextLine();

                    System.out.print("Name: ");
                    String name = scanner.nextLine();

                    System.out.print("Age: ");
                    int age = scanner.nextInt();

                    service.addStudent(new Student(id, name, age));
                    break;

                case 2:
                    service.getAllStudents().forEach(System.out::println);
                    break;

                case 3:
                    System.out.print("Enter ID: ");
                    int searchId = scanner.nextInt();
                    System.out.println(service.findStudentById(searchId));
                    break;

                case 4:
                    return;
            }
        }
    }
}
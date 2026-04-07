using System;
using TaskManager.Services;

class Program
{
    static void Main()
    {
        var service = new TaskService();

        while (true)
        {
            Console.WriteLine("\n1. Add Task");
            Console.WriteLine("2. View Tasks");
            Console.WriteLine("3. Complete Task");
            Console.WriteLine("4. Delete Task");
            Console.WriteLine("5. Exit");

            Console.Write("Choose: ");
            int choice = int.Parse(Console.ReadLine());

            switch (choice)
            {
                case 1:
                    Console.Write("Enter task: ");
                    string name = Console.ReadLine();
                    service.AddTask(name);
                    break;

                case 2:
                    service.GetTasks().ForEach(Console.WriteLine);
                    break;

                case 3:
                    Console.Write("Enter ID: ");
                    service.CompleteTask(int.Parse(Console.ReadLine()));
                    break;

                case 4:
                    Console.Write("Enter ID: ");
                    service.DeleteTask(int.Parse(Console.ReadLine()));
                    break;

                case 5:
                    return;
            }
        }
    }
}
using TaskManager.Models;
using System.Collections.Generic;
using System.Linq;

namespace TaskManager.Services
{
    public class TaskService
    {
        private List<TaskItem> tasks = new();
        private int counter = 1;

        public void AddTask(string name)
        {
            tasks.Add(new TaskItem
            {
                Id = counter++,
                Name = name,
                IsCompleted = false
            });
        }

        public List<TaskItem> GetTasks()
        {
            return tasks;
        }

        public void CompleteTask(int id)
        {
            var task = tasks.FirstOrDefault(t => t.Id == id);
            if (task != null) task.IsCompleted = true;
        }

        public void DeleteTask(int id)
        {
            tasks.RemoveAll(t => t.Id == id);
        }
    }
}
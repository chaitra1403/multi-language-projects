namespace TaskManager.Models
{
    public class TaskItem
    {
        public int Id { get; set; }
        public string Name { get; set; }
        public bool IsCompleted { get; set; }

        public override string ToString()
        {
            return $"[{Id}] {Name} - {(IsCompleted ? "Done" : "Pending")}";
        }
    }
}
import React from 'react';
import TaskItem from './TaskItem';

interface Task {
  id: number;
  title: string;
  description: string;
  due_date: string;
  status: number;
  created_at: string;
  updated_at: string;
}

interface TaskListProps {
  tasks: Task[];
  fetchTasks: () => void;
}

const TaskList: React.FC<TaskListProps> = ({ tasks, fetchTasks }) => {
  return (
    <ul>
      {tasks.map(task => (
        <TaskItem key={task.id} task={task} fetchTasks={fetchTasks} />
      ))}
    </ul>
  );
};

export default TaskList;

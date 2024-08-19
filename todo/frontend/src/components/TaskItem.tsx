import React from 'react';

interface Task {
  id: number;
  title: string;
  description: string;
  due_date: string;
  status: number;
  created_at: string;
  updated_at: string;
}

interface TaskItemProps {
  task: Task;
  fetchTasks: () => void;
}

const TaskItem: React.FC<TaskItemProps> = ({ task, fetchTasks }) => {
  const updateStatus = async (status: number) => {
    try {
      const response = await fetch(`${process.env.REACT_APP_API_URL}/tasks/${task.id}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ ...task, status }),
      });
      if (response.ok) {
        fetchTasks(); // タスクリストを更新
      }
    } catch (error) {
      console.error('Error updating task status:', error);
    }
  };

  return (
    <li>
      <span>{task.title} - {task.description}</span>
      <button onClick={() => updateStatus(task.status === 0 ? 1 : 0)}>
        {task.status === 0 ? 'Mark as Done' : 'Mark as Undone'}
      </button>
    </li>
  );
};

export default TaskItem;

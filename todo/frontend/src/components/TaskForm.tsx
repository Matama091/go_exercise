import React, { useState } from 'react';

interface TaskFormProps {
  onTaskAdded: () => void;
}

const TaskForm: React.FC<TaskFormProps> = ({ onTaskAdded }) => {
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');

  const addTask = async (e: React.FormEvent) => {
    e.preventDefault();
    await fetch('http://localhost:8080/tasks', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ title, description, due_date: '2023-12-31', status: 0 })
    });
    onTaskAdded();
    setTitle('');
    setDescription('');
  };

  return (
    <form onSubmit={addTask}>
      <input value={title} onChange={e => setTitle(e.target.value)} placeholder="Title" />
      <input value={description} onChange={e => setDescription(e.target.value)} placeholder="Description" />
      <button type="submit">Add Task</button>
    </form>
  );
};

export default TaskForm;

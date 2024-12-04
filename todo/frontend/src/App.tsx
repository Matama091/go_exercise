import React, { useState, useEffect } from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import TaskList from './components/TaskList';
import TaskForm from './components/TaskForm';
import TaskItem from './components/TaskItem';

interface Task {
  id: number;
  title: string;
  description: string;
  due_date: string;
  status: number;
  created_at: string;
  updated_at: string;
}

const App: React.FC = () => {
  const [tasks, setTasks] = useState<Task[]>([]);
  const [selectedTask, setSelectedTask] = useState<Task | null>(null);

  const fetchTasks = async () => {
    // ここにタスクを取得するためのロジックを追加します
    // 例: APIからタスクを取得してsetTasksを呼び出す
    const response = await fetch('/api/tasks');
    const data = await response.json();
    setTasks(data);
  };

  useEffect(() => {
    fetchTasks();
  }, []);

  const handleAddTask = (task: { title: string; description: string; due_date: string; status: number }) => {
    console.log('Task added:', task);
    // ここにタスクを追加するためのロジックを追加できます
  };

  const handleSelectTask = (task: Task) => {
    setSelectedTask(task);
  };

  return (
    <BrowserRouter>
      <p>
        Edit <code>src/App.tsx</code> and save to reload.
      </p>
      <Routes>
        <Route path="/" element={<TaskList tasks={tasks} fetchTasks={fetchTasks} onSelectTask={handleSelectTask} />} />
        <Route path="/tasks" element={<TaskList tasks={tasks} fetchTasks={fetchTasks} onSelectTask={handleSelectTask} />} />
        <Route path="/task-form" element={<TaskForm onAddTask={handleAddTask} />} />
        <Route path="/task-item" element={selectedTask && <TaskItem task={selectedTask} fetchTasks={fetchTasks} />} />
      </Routes>
    </BrowserRouter>
  );
};

export default App;

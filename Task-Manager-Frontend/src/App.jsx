import React, { useState, useEffect } from 'react';
import TaskInput from './components/TaskInput';
import TaskList from './components/TaskList';

const App = () => {
  const [tasks, setTasks] = useState([]);
  const [title, setTitle] = useState('');
  const [error, setError] = useState(null);
  const [loading, setLoading] = useState(true);

  const fetchTasks = async () => {
    setLoading(true);
    try {
      const res = await fetch('http://backend:8080/tasks');
      if (!res.ok) throw new Error('Failed to fetch tasks');
      const data = await res.json();
      console.log('Tasks data after fetch:', data); // Debug log
      setTasks(data || []);
      setError(null);
    } catch (err) {
      setError(err.message);
      setTasks([]); // Set to empty array on error
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchTasks();
  }, []);

  const addTask = async () => {
    if (!title.trim()) return;
    try {
      const res = await fetch('http://backend:8080/tasks', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ title, completed: false }),
      });
      if (!res.ok) throw new Error('Failed to add task');
      setTitle('');
      fetchTasks();
    } catch (err) {
      console.error('Error adding task:', err);
    }
  };

  const toggleTask = async (id, completed) => {
    // Optimistically update the local state, preserving title
    const updatedTasks = tasks.map(task =>
      task.id === id ? { ...task, completed: !completed } : task
    );
    console.log('Tasks after optimistic update:', updatedTasks); // Debug log
    setTasks(updatedTasks);
    try {
      const res = await fetch(`http://backend:8080/tasks/${id}`, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ completed: !completed }), // Send only completed
      });
      if (!res.ok) throw new Error('Failed to update task');
      fetchTasks(); // Sync with backend
    } catch (err) {
      console.error('Error toggling task:', err);
      // Revert on error
      fetchTasks();
    }
  };

  const deleteTask = async (id) => {
    setTasks(tasks.filter(task => task.id !== id));
    try {
      const res = await fetch(`http://backend:8080/tasks/${id}`, {
        method: 'DELETE',
      });
      if (!res.ok) throw new Error('Failed to delete task');
      fetchTasks();
    } catch (err) {
      console.error('Error deleting task:', err);
      fetchTasks();
    }
  };

  return (
    <div>
      <div className="max-w-2xl mx-auto p-6">
        <h1 className="text-3xl font-bold text-center mb-6">Task Manager</h1>
        <TaskInput title={title} setTitle={setTitle} addTask={addTask} />
        <TaskList 
          tasks={tasks} 
          toggleTask={toggleTask} 
          deleteTask={deleteTask} 
          error={error} 
          loading={loading} 
        />
      </div>
    </div>
  );
};

export default App;
import React from 'react';
import TaskItem from './TaskItem';

const TaskList = ({ tasks, toggleTask, deleteTask, error, loading }) => {
  return (
    <div>
      {loading && <p className="text-gray-500 mb-4">Loading tasks...</p>}
      {error && <p className="text-red-500 mb-4">{error}</p>}
      <ul className="space-y-4">
        {(tasks || []).map((task) => (
          <TaskItem key={task.id} task={task} toggleTask={toggleTask} deleteTask={deleteTask} />
        ))}
      </ul>
    </div>
  );
};

export default TaskList;
import React from 'react';

const TaskItem = ({ task, toggleTask, deleteTask }) => {
  return (
    <li
      key={task.id}
      className="flex items-center justify-between p-4 bg-gray-100 rounded-md"
    >
      <div className="flex items-center gap-3">
        <input
          type="checkbox"
          checked={task.completed}
          onChange={() => toggleTask(task.id, task.completed)}
          className="h-5 w-5 text-blue-600"
        />
        <span
          className={`${
            task.completed ? 'line-through text-gray-500' : 'text-gray-800'
          }`}
        >
          {task.title} {/* Ensure only the original title is shown */}
        </span>
      </div>
      <button
        onClick={() => deleteTask(task.id)}
        className="text-red-500 hover:text-red-700 transition"
      >
        Delete
      </button>
    </li>
  );
};

export default TaskItem;
import React from 'react'

const TaskInput = ({ title, setTitle, addTask }) => {
  return (
    <div className='flex gap-4 mb-6'>
        <input
            type='text'
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            placeholder='Add a new task'
            className='flex-1 p-2 border border-gray-300 rounded-md focus:outline-none focus:ring-blue-500'
        />
        <button
            onClick={addTask} 
            className='bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700 transition'
        >
            Add
        </button>
    </div>
  )
}

export default TaskInput
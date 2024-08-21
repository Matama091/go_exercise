import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './App';


const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);
root.render(
  <>
    <h1>Task Manager</h1>
    <div className='container'>
        <App />
    </div>
  </>
);

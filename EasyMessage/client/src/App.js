import React from 'react';
import { Route, Routes } from 'react-router';
import { Auth, Home } from './pages';

function App() {
  return (
    <div className="wrapper">
      <Routes>
        <Route path="/" element={<Auth />} />
        <Route path="/login" element={<Auth />} />
        <Route path="/register" element={<Auth />} />
        <Route path="/im" element={<Home />} />
      </Routes>
    </div>
  );
}

export default App;

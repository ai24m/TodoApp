import React from 'react';
import './App.css';
import { Header } from './components/Header';
import { TodoList } from './components/TodoList';



function App() {
  return (
    <div className="App">
      <Header/> 
      <TodoList todos={[
        {title: "Get Cat", description: "Today", isCompleted: true},
        {title: "Pet Cat", description: "Tomorrow", isCompleted: false},
        {title: "Feed Cat", description: "Tomorrow", isCompleted: false}
        ]}></TodoList>

    </div>
  );
}

export default App;

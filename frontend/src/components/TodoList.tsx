import React from "react"
import { Todo } from "../entities/Todo"
import { TodoItem } from "./TodoItem"
import './TodoList.scss'

type Props = {
    todos: Todo[];
}

export const TodoList: React.FC<Props> = ({ todos }) => {
    return (
        <ul className = "todo-list">
            {
                todos.map((todo , i)=> (
                    // <li key={i}> {todo.title}  </li>
                    <li key={i}> 
                        <TodoItem 
                        title={todo.title} 
                        description={todo.description} 
                        isCompleted={todo.isCompleted}></TodoItem>
                    </li>  //replace with created TodoItem component that calls Todo entity
                ))
            }
        </ul>
    )
}
import { Todo } from "../entities/Todo";
import React from "react";
import './TodoItem.scss';

export const TodoItem: React.FC<Todo> = ({title,description,isCompleted}) => {
    return (
        <article className="todo-item">
            <section className="todo-item-text">
                <h3>{title}</h3>
                <p>{description}</p>
                <input type="checkbox" defaultChecked={isCompleted}/>
            </section>
        </article>
    )
}
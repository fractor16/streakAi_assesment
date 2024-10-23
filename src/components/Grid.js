import React, { useState, useEffect } from 'react';
import axios from 'axios';
import './Grid.css'; 

const Grid = () => {
    const [start, setStart] = useState(null);
    const [end, setEnd] = useState(null);
    const [path, setPath] = useState([]);
    
    
    useEffect(() => {
        if (start && end) {
            const fetchPath = async () => {
                try {
                    const response = await axios.post('http://localhost:8080/find-path', {
                        start: { x: start[0], y: start[1] }, 
                        end: { x: end[0], y: end[1] }
                    });
                    console.log("Path from backend:", response.data.path);
                    setPath(response.data.path);
                } catch (error) {
                    console.error('Error fetching path:', error);
                    alert('No path found!');
                }
            };

            
            fetchPath();
        }
    }, [start, end]);  

    const handleClick = (row, col) => {
        if (!start) {
            setStart([row, col]);
            console.log("Start point set:", [row, col]);  
        } else if (!end) {
            setEnd([row, col]);
            console.log("End point set:", [row, col]);  
        }
    };

    const isStart = (row, col) => start && start[0] === row && start[1] === col;
    const isEnd = (row, col) => end && end[0] === row && end[1] === col;
    const isPath = (row, col) => {
        return path.some(p => p.x === row && p.y === col);  
    };
    
    const resetGrid = () => {
        setStart(null);
        setEnd(null);
        setPath([]);
    };

    const renderGrid = () => {
        const rows = [];
        for (let i = 0; i < 20; i++) {
            const cols = [];
            for (let j = 0; j < 20; j++) {
                cols.push(
                    <div
                        key={`${i}-${j}`}
                        className={`cell ${isStart(i, j) ? 'start' : ''} ${isEnd(i, j) ? 'end' : ''} ${isPath(i, j) ? 'path' : ''}`}
                        onClick={() => handleClick(i, j)}
                    />
                );
            }
            rows.push(<div key={i} className="row">{cols}</div>);
        }
        return rows;
    };

    return (
        <div className="grid-container">
            <button onClick={resetGrid}>Reset Grid</button>
            <div className="grid">
                {renderGrid()}
            </div>
        </div>
    );
};

export default Grid;

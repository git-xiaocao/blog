import React from 'react';

import './App.css';
import API from "./api/api";

function App() {
    return (
        <div className="App">
            <button onClick={() => {
                API.Article.GetList(0, 30, null).then(result => {
                    console.log(result);
                })
            }}>

            </button>
        </div>
    );
}

export default App;

import React, { useState, useEffect } from 'react';
import CacheContents from './CacheContents';
import SetKeyValueForm from './SetKeyValueForm';
import GetCacheByKey from './GetCacheByKey'; 
import './App.css'; 

function App() {
    const [cacheContents, setCacheContents] = useState([]);

    useEffect(() => {
        // Fetch cache contents from backend
        fetchCacheContents();
    }, []);

    const fetchCacheContents = () => {
        fetch('http://localhost:8080/cache')
            .then(response => response.json())
            .then(data => setCacheContents(data))
            .catch(error => console.error('Error fetching cache contents:', error));
    };

    return (
        <div className="App">
            <h1>LRU Cache Viewer</h1>
            <CacheContents cacheContents={cacheContents} />
            <SetKeyValueForm onSet={() => fetchCacheContents()} />
            <GetCacheByKey /> 
        </div>
    );
}

export default App;

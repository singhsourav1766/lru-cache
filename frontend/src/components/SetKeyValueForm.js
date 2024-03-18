import React, { useState } from 'react';
import './SetKeyValueForm.css';

function SetKeyValueForm({ onSet }) {
    const [key, setKey] = useState('');
    const [value, setValue] = useState('');

    const handleSubmit = (e) => {
        e.preventDefault();
        // Send key-value pair to backend to set in cache
        fetch('http://localhost:8080/cache/' + key, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(value),
        })
            .then(() => {
                setKey('');
                setValue('');
                onSet();
            })
            .catch(error => console.error('Error setting key-value pair:', error));
    };

    return (
        <div>
            <h2>Set Key-Value Pair</h2>
            <form onSubmit={handleSubmit}>
                <input type="text" placeholder="Key" value={key} onChange={(e) => setKey(e.target.value)} />
                <input type="text" placeholder="Value" value={value} onChange={(e) => setValue(e.target.value)} />
                <button type="submit">Set</button>
            </form>
        </div>
    );
}

export default SetKeyValueForm;

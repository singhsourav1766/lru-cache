import React, { useState } from 'react';

function GetCacheByKey() {
    const [key, setKey] = useState('');
    const [value, setValue] = useState('');
    const [error, setError] = useState(null);

    const handleChange = (event) => {
        setKey(event.target.value);
    };

    const handleSubmit = (event) => {
        event.preventDefault();
        setError(null);

        fetch(`http://localhost:8080/cache/${key}`)
            .then(response => {
                if (!response.ok) {
                    throw new Error('Failed to fetch cache item');
                }
                return response.json();
            })
            .then(data => {
                setValue(data);
            })
            .catch(error => {
                setError(error.message);
            });
    };

    return (
        <div>
            <h2>Get Cache by Key</h2>
            <form onSubmit={handleSubmit}>
                <label>
                    Key:
                    <input type="text" value={key} onChange={handleChange} />
                </label>
                <button type="submit">Fetch</button>
            </form>
            {value && (
                <div>
                    <h3>Cache Value:</h3>
                    <p>{value}</p>
                </div>
            )}
            {error && <p>Error: {error}</p>}
        </div>
    );
}

export default GetCacheByKey;

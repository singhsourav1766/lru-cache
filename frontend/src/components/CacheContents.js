import React from 'react';
function CacheContents({ cacheContents }) {
    return (
        <div>
            <h2>Cache Contents</h2>
            <ul>
                {cacheContents.map(entry => (
                    <li key={entry.Key}>
                        <strong>{entry.Key}:</strong> {entry.Value}
                    </li>
                ))}
            </ul>
        </div>
    );
}

export default CacheContents;

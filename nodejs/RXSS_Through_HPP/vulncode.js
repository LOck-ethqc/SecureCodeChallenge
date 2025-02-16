const express = require('express');
const app = express();
const port = 3001;

app.use(express.json());

function sanitizeName(name) {
    const blacklist = [
        '@', '#', '$', '%', '^', '&', '*', '(', ')',
        '-', '=', '+', '[', ']', '{', '}', ';', ':',
        "'", '"', '<', '>', ',', '.', '?', '/', '\\', '|'
    ];
    return !blacklist.some(char => name.includes(char));
}

app.get('/', (req, res) => {
    let name = req.query.name;

    // Default name if none provided
    if (!name) {
        name = 'Guest';
    }

    // Check if the name is valid based on the blacklist
    if (!sanitizeName(name)) {
        return res.status(400).json({ error: 'Name must only contain letters (a-z, A-Z)' });
    }

    // Send greeting message
    res.send(`<h1>Hello, ${name}!</h1>`);
});

app.listen(port, '0.0.0.0', () => {
    console.log(`Server running on http://0.0.0.0:${port}`);
});


import logo from './logo.svg';
import './App.css';
import React, {useEffect, useState} from 'react';

function App() {
  const [messages, setMessages] = useState([]);

  useEffect(() => {
    fetch("http://localhost:8082/api/messages")
      .then(response => response.json())
      .then(data => setMessages(data))
      .catch(err => console.error(err));
  }, []);


  return (
    <div className="App">
      <h1>Messages</h1>
      {messages.length ? (
        <ul>
          {messages.map(message => (
            <li key={message.id}>{message.content}</li>
          ))}
        </ul>
      ) : (
        <p>No messages available</p>
      )}
    </div>
  );

}

export default App;

import { useState } from "react";
import './InputComponent.css'

export default function InputComponent() {
    const [text, setText] = useState('');

    const handleTextChange = (ev) => {
        setText(ev.target.value);
    }

    const handleTextSubmit = (ev) => {
        ev.preventDefault();
        console.log('submitted!!! ', text);
    }

    return (
        <form onSubmit={handleTextSubmit}>
            <label>What did you do?</label>
            <input id="text-input" type="input/text" value={text} onChange={handleTextChange} placeholder="Type here"></input>
            <button type='submit'>Submit</button>
        </form>
    );
}


import { useState } from "react";
import './InputComponent.css'

export default function InputComponent({ onSubmit, loading }) {
    const ZODIAC_SIGNS = ['Aries', 'Taurus', 'Gemini', 'Cancer', 'Leo', 'Virgo', 'Libra', 'Scorpio', 'Sagittarius', 'Capricorn', 'Aquarius', 'Pisces']

    const [sign, setSign] = useState('');

    const [text, setText] = useState('');

    const handleTextChange = (ev) => {
        setText(ev.target.value);
    }

    const handleTextSubmit = (ev) => {
        ev.preventDefault();
        if (text.trim()) onSubmit(text, sign);
    }

    return (
        <form className="excuse-form" onSubmit={handleTextSubmit}>
            <label>What did you do?</label>
            <textarea id="text-input" value={text} onChange={handleTextChange} placeholder="Type here" />
            <select value={sign} required onChange={(e) => setSign(e.target.value)}> 
                <option value="" required>Select your Zodiac sign</option>
                {ZODIAC_SIGNS.map((s) => <option key={s} value={s}>{s}</option>)}
            </select>
            <button type='submit' disabled={loading}>{loading ? 'Hmmm...' : 'Submit'}</button>
        </form>
    );
}


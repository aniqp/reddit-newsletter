import React, { useState } from 'react';
import { addUserEmail } from '@/db';
import { useDispatch } from 'react-redux';
import { setUser } from '@/redux/slice'

const SignUpModal = ({ open, userId }) => {
    const dispatch = useDispatch()
    const [email, setEmail] = useState('');
    const [submitted, setSubmitted] = useState(false);

    // Handle email input change
    const handleEmailChange = (e) => {
        setEmail(e.target.value);
    };

    // Handle form submission
    const handleSubmit = (e) => {
        e.preventDefault();
        if (email) {
            setSubmitted(true);
            dispatch(setUser({reddit: userId, email: email}))
            addUserEmail(userId, email)
            console.log("Email submitted: ", email);
        }
    };

    if (!open || submitted) return null;

    return (
        <div style={{
            position: 'fixed',
            top: 0,
            left: 0,
            width: '100%',
            height: '100%',
            backgroundColor: 'rgba(0, 0, 0, 0.5)',
            display: 'flex',
            justifyContent: 'center',
            alignItems: 'center',
            zIndex: 1000,
        }}>
            <div style={{
                backgroundColor: '#fff',
                padding: '20px',
                borderRadius: '5px',
                boxShadow: '0 4px 6px rgba(0, 0, 0, 0.1)',
                width: '300px',
            }}>
                <form onSubmit={handleSubmit}>
                    <div style={{ marginBottom: '15px' }}>
                        <label htmlFor="email" style={{ display: 'block', marginBottom: '5px' }}>Email:</label>
                        <input
                            type="email"
                            id="email"
                            value={email}
                            onChange={handleEmailChange}
                            required
                            placeholder="Enter your email"
                            style={{
                                width: '100%',
                                padding: '8px',
                                borderRadius: '4px',
                                border: '1px solid #ccc',
                            }}
                        />
                    </div>
                    <button
                        type="submit"
                        disabled={!email}
                        style={{
                            padding: '10px 15px',
                            border: 'none',
                            borderRadius: '4px',
                            backgroundColor: '#007bff',
                            color: '#fff',
                            cursor: 'pointer',
                            width: '100%',
                        }}
                    >
                        Submit
                    </button>
                </form>
            </div>
        </div>
    );
}

export default SignUpModal;

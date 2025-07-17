import { useState } from 'react';
import '../styles/createIssue.css'

function CreateForm() {

    const [summary, setSummary] = useState('');
    const [description, setDescription] = useState('');

    const handleCreate = async (e) => {
        e.preventDefault();

        const data = {
            summary: summary,
            description: description,
        };

        try {
            const responce = await fetch(`/api/tasks`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(data),
            });
            const result = await responce.json();
            console.log('Responce from server:', result)
        } catch (error) {
            console.error("Error while post", error)
        }
        window.location.reload()
    };

    return (
        <center><form onSubmit={handleCreate} className='form'>
            <div className='input-label'>
            <label>
                Summary
                <input type="text" name="summary" className="label-input" onChange={e => setSummary(e.target.value)}/>
            </label>
            </div>

            <div className='input-label'>
            <label>
                Description
                <input type="text" name="description" className="label-input" onChange={e => setDescription(e.target.value)}/>
            </label>
            </div>

            <div className='create-btn'>
            <button>Create!</button>
            </div>
        </form></center>
    );
}

export default CreateForm;
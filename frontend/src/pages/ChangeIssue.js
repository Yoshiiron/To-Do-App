import { useEffect, useState } from 'react';
import '../styles/createIssue.css'

function ChangeForm({issue}) {

    
    const [summary, setSummary] = useState('');
    const [description, setDescription] = useState('');
    const [status, setStatus] = useState('To Do')
    const statuses = ['To Do', 'In Progress', 'Done']

    useEffect(() => {
        if (issue) {
            setSummary(issue.Summary || '');
            setDescription(issue.Description || '');
            setStatus(issue.Status || '');
        }
    }, [issue])

    const handleChange = async (e) => {
        e.preventDefault();

        const data = {
            summary: summary,
            description: description,
            status: status,
        };

        try {
            const responce = await fetch(`http://localhost:8080/task/${issue.IssueID}`, {
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
        <center><form onSubmit={handleChange} className='form'>
            <div className='input-label'>
            <label>
                Summary
                <input type="text" name="summary" className="label-input" value={summary} onChange={e => setSummary(e.target.value)}/>
            </label>
            </div>

            <div className='input-label'>
            <label>
                Description
                <input type="text" name="description" className="label-input" value={description} onChange={e => setDescription(e.target.value)}/>
            </label>
            </div>

            <div className='input-label'>
            <label>
                <div>Status</div>
                <select value={status} onChange={(e) => setStatus(e.target.value)}>
                {statuses.map((s) => (
                    <option key={s} value={s}>
                        {s}
                    </option>
                ))}
                </select>
            </label>
            </div>

            <div className='create-btn'>
            <button>Change!</button>
            </div>
        </form></center>
    );
}

export default ChangeForm;
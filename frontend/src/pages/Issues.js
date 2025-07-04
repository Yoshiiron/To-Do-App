import { useEffect, useState } from "react";
import { Helmet } from 'react-helmet';
import axios from 'axios';

import '../styles/Issues.css';
import Modal from "./modal";
import ChangeForm from "./ChangeIssue";

function ItemList() {
const [issues, setIssues] = useState([]);
const [filter, setFilter] = useState('All');
const [modalOpen, setModalOpen] = useState(false);
const [idToEdit, setIDToEdit] = useState(null);

useEffect(() => {
  axios.get("http://localhost:8080/tasks")
  .then(resp => {
    setIssues(resp.data.response);
  })
}, [])

const deleteIssue = (id) => {
  axios.delete(`http://localhost:8080/tasks/${id}`)
  .then(() => {
    setIssues(prevIssue => prevIssue.filter(item => item.IssueID !== id));
  })
  .catch(error => {
    console.error('Ошибка при удалении:', error)
  });
};

const filterIssues = filter === 'All'
? issues
: issues.filter(issues => issues.Status === filter)

  return (
    <>
      <Helmet>
        <title>ToDo</title>
      </Helmet>
      <main>
        <div>
          <center><h2>Issues</h2></center>

          <div className="filter-buttons">
            {['All', 'To Do', 'In Progress', 'Done'].map(status => (
              <button
              key={status}
              className={filter === status ? 'active-filter' : ''}
              onClick={() => setFilter(status)}
              >
                {status}
              </button>
            ))}
          </div>

          <ul class="issues">
            {filterIssues.map((issue, index) => (
              <li key={index} className="issue-card">
                <p>Summary: {issue.Summary}</p>
                <p>Description: {issue.Description}</p>
                <p>Status: {issue.Status}</p>
                <p className="issue-id">ID: {issue.IssueID}</p>
                <button className="delete-issue" onClick={() => deleteIssue(issue.IssueID)}>❌</button>
                <button className="change-issue" onClick={() => {setModalOpen(true); setIDToEdit(issue);}}>🖋</button>

              </li>
            ))}
          </ul>

          <Modal isOpen={modalOpen} onClose={() => setModalOpen(false)}>
            <center><h2>Change an Issue</h2></center>
            {idToEdit && <ChangeForm issue={idToEdit}/>}
          </Modal>
        </div>
      </main>
    </>
  );
}

export default ItemList;
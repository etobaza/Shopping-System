import React, { useState, useEffect } from 'react';
import * as userService from "../services/user";
import Navbar from "./Navbar";

const AdminPanel = () => {
    const [users, setUsers] = useState([]);
    const [search, setSearch] = useState('');
    const [sortOrder, setSortOrder] = useState('none');

    useEffect(() => {
        userService.fetchUsers().then((data) => setUsers(data));
    }, []);

    const handleDeleteUser = async (id) => {
        try {
            await userService.deleteUser(id);
            const updatedUsers = users.filter((user) => user.id !== id);
            setUsers(updatedUsers);
        } catch (error) {
            console.error("Error deleting user:", error);
        }
    };

    const filteredUsers = users.filter((user) => user.username.toLowerCase().includes(search.toLowerCase()) && user.usertype !== "admin");
    const sortedUsers = filteredUsers.slice().sort((a, b) => {
        if (sortOrder === 'increasing') {
            return a.id - b.id;
        } else if (sortOrder === 'decreasing') {
            return b.id - a.id;
        } else {
            return 0;
        }
    });

    return (
        <div>
            <Navbar />
            <h1>Admin Panel</h1>
            <div>
                <input
                    type="text"
                    placeholder="Search by username"
                    value={search}
                    onChange={(e) => setSearch(e.target.value)}
                />
                <select value={sortOrder} onChange={(e) => setSortOrder(e.target.value)}>
                    <option value="none">No Sort</option>
                    <option value="increasing">Sort Increasing</option>
                    <option value="decreasing">Sort Decreasing</option>
                </select>
            </div>
            <table>
                <thead>
                <tr>
                    <th>ID</th>
                    <th>Username</th>
                    <th>Email</th>
                    <th>First Name</th>
                    <th>Last Name</th>
                    <th>Address</th>
                    <th>Phone</th>
                    <th>User Type</th>
                    <th>Action</th>
                </tr>
                </thead>
                <tbody>
                {sortedUsers.map((user) => (
                    <tr key={user.id}>
                        <td>{user.id}</td>
                        <td>{user.username}</td>
                        <td>{user.email}</td>
                        <td>{user.firstname}</td>
                        <td>{user.lastname}</td>
                        <td>{user.address}</td>
                        <td>{user.phone}</td>
                        <td>{user.usertype}</td>
                        <td>
                            <button onClick={() => handleDeleteUser(user.id)}>Delete</button>
                        </td>
                    </tr>
                ))}
                </tbody>
            </table>
        </div>
    );
};

export default AdminPanel;

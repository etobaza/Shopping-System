import {API_BASE_URL} from "../api/config";

const register = async (userData) => {
    try {
        const response = await fetch(`${API_BASE_URL}/register`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(userData),
        });

        if (!response.ok) {
            throw new Error("Error registering user");
        }

        const data = await response.json();
        return { success: true, data };
    } catch (error) {
        console.error("Error registering user:", error);
        return { success: false, message: "Registration failed. Please try again." };
    }
};

const login = async (credentials) => {
    try {
        const response = await fetch(`${API_BASE_URL}/login`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(credentials),
        });

        if (!response.ok) {
            throw new Error("Error logging in user");
        }

        const data = await response.json();
        localStorage.setItem('userId', data.id);
        localStorage.setItem('userType', data.usertype);
        return { success: true, data };
    } catch (error) {
        return { success: false, message: "Incorrect username or password." };
    }
};

export const fetchUsers = async () => {
    try {
        const response = await fetch(`${API_BASE_URL}/users`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
            },
        });

        if (!response.ok) {
            throw new Error('Error fetching users');
        }

        const data = await response.json();
        return data;
    } catch (error) {
        console.error('Error fetching users:', error);
    }
};


export const fetchUser = async (id) => {
    try {
        const response = await fetch(`${API_BASE_URL}/users/${id}`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
            },
        });

        if (!response.ok) {
            throw new Error('Error fetching user data');
        }

        const data = await response.json();
        console.log('User data:', data);
        return data;
    } catch (error) {
        console.error('Error fetching user data:', error);
    }
};

export const updateUser = async (id, userData) => {
    try {
        const response = await fetch(`${API_BASE_URL}/users/${id}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(userData),
        });

        if (!response.ok) {
            throw new Error('Error updating user');
        }

        const data = await response.json();
        return { success: true, data };
    } catch (error) {
        console.error('Error updating user:', error);
        return { success: false, message: 'Update failed. Please try again.' };
    }
};

export const deleteUser = async (id) => {
    try {
        const response = await fetch(`${API_BASE_URL}/users/${id}`, {
            method: 'DELETE',
        });

        if (!response.ok) {
            throw new Error('Error deleting user');
        }
    } catch (error) {
        console.error('Error deleting user:', error);
    }
};

export const handleLogout = async () => {
    try {
        const response = await fetch('/logout', { method: 'POST' });
        if (!response.ok) {
            throw new Error('Failed to logout');
        }
        return true;
    } catch (error) {
        console.error(error);
        return false;
    }
};

export default {
    register,
    login,
    handleLogout,
    fetchUsers,
    fetchUser,
    deleteUser,
};

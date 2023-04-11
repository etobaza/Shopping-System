import {API_BASE_URL} from "../api/config";

export const fetchItems = async () => {
    try {
        const response = await fetch(`${API_BASE_URL}/items`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
            },
        });

        if (!response.ok) {
            throw new Error('Error fetching items');
        }

        const data = await response.json();
        return data;
    } catch (error) {
        console.error('Error fetching items:', error);
    }
};

export const fetchItemsSeller = async () => {
    const userID = localStorage.getItem('userId');

    try {
        const response = await fetch(`${API_BASE_URL}/shop/items?user_id=${userID}`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
            },
        });

        if (!response.ok) {
            throw new Error('Error fetching items');
        }

        const data = await response.json();
        console.log('Data:', data);
        return data;
    } catch (error) {
        console.error('Error fetching items:', error);
    }
};

export const createItem = async (itemData) => {
    try {
        const response = await fetch(`${API_BASE_URL}/shop/items`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(itemData),
        });

        if (!response.ok) {
            throw new Error('Error creating item');
        }

        const data = await response.json();
        return data;
    } catch (error) {
        console.error('Error creating item:', error);
    }
};

export const updateItem = async (id, itemData) => {
    try {
        const response = await fetch(`${API_BASE_URL}/items/${id}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(itemData),
        });

        if (!response.ok) {
            throw new Error('Error updating item');
        }

        const data = await response.json();
        return { success: true, data };
    } catch (error) {
        console.error('Error updating item:', error);
        return { success: false, message: 'Update failed. Please try again.' };
    }
};

export const deleteItem = async (id) => {
    try {
        const response = await fetch(`${API_BASE_URL}shop/items/${id}`, {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/json',
            },
        });

        if (!response.ok) {
            throw new Error('Error deleting item');
        }

        const data = await response.json();
        console.log('Data:', data);
        return data;
    } catch (error) {
        console.error('Error deleting item:', error);
    }
};

export const parseCategories = (items) => {
    const categories = new Set();
    items.forEach(item => {
        categories.add(item.category);
    });
    return Array.from(categories);
}

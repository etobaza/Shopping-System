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

export const parseCategories = (items) => {
    const categories = new Set();
    items.forEach(item => {
        categories.add(item.category);
    });
    return Array.from(categories);
}

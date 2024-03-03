import React from 'react';
import { Dialog as BaseDialog } from 'components';
import { DialogItem } from "components";
const Dialog = ({items, search}) => {
    let filteredItems = items;

    search="Але";
    if (!items || items.length === 0) {
        return <div className="dialogs">No dialogs found</div>;
    }
    if (search) {
        filteredItems = items.filter(item => item.user.name.toLowerCase().includes(search.toLowerCase()));
    }

    return (
        <div className="dialogs">
            {items.map((filteredItems, index) => (
                <DialogItem key={index} user={filteredItems.user} message={filteredItems.message}/>
            ))}
        </div>
    );
};

export default Dialog;
import React from 'react';
import { DialogItem } from "components";

const Dialog = ({ items }) => {
    if (!items || items.length === 0) {
        return <div className="dialogs">No dialogs found</div>;
    }

    return (
        <div className="dialogs">
            {items.map((item, index) => (
                <DialogItem key={index} user={item.user} message={item.message} />
            ))}
        </div>
    );
};

export default Dialog;

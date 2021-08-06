import React, { useState } from "react";
import Form from "react-bootstrap/Form";
import InputGroup from "react-bootstrap/InputGroup";
import Button from "react-bootstrap/Button";

interface IAddEntry {
  onAdd: (val: string) => void;
  label: string;
  placeholder?: string;
}

export const AddEntry: React.FC<IAddEntry> = (props: IAddEntry) => {
  const [value, setValue] = useState("");
  const onSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    props.onAdd(value);
  };

  return (
    <Form onSubmit={onSubmit}>
      <Form.Group>
        <Form.Label>{props.label}</Form.Label>
        <InputGroup hasValidation>
          <Form.Control
            required
            placeholder={props.placeholder}
            value={value}
            onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
              setValue(e.currentTarget.value)
            }
          />
          <Button variant="primary" type="submit">
            Add
          </Button>
        </InputGroup>
      </Form.Group>
    </Form>
  );
};

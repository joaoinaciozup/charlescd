export const conditionOptions = [
  { value: 'EQUAL', label: 'Equal' },
  { value: 'GREATER_THAN', label: 'Higher than' },
  { value: 'LOWER_THAN', label: 'Lower than' }
];

export const operatorsOptions = [
  { value: '=', label: 'Equal' },
  { value: '!=', label: 'Different' },
  { value: '!~', label: 'Regex' }
];

export const FILTER = {
  field: '',
  condition: '',
  value: ''
};

export const defaultFilterValues = [FILTER];
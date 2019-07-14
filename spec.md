gantt-taro table spec
===

## users
- id
- name

## projects
- id
- name
- description
- is_private

## project_members
- id
- project_id
- user_id

## tasks
- id
- name
- user_id
- project_id
- parent_task_id
- started_at
- ended_at
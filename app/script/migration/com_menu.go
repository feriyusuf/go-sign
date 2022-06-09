package migration

const SqlDefaultMenusInit = `insert into public.com_menu (id, name, label, icon, is_active, parent_id, type, deleted_at, ui_path, api_path, created_at, updated_at, created_by, updated_by, deleted_by)
values  (1, 'MENU_MAIN', 'Menu', null, true, null, 'SCREEN', null, null, null, '2022-06-09 16:15:52.394143 +00:00', '2022-06-09 16:15:52.394143 +00:00', null, null, null),
        (2, 'MENU_ADD', 'Add', null, true, 1, 'ACTION', null, null, null, '2022-06-09 16:24:20.328077 +00:00', '2022-06-09 16:24:20.328077 +00:00', null, null, null),
        (3, 'MENU_DETAIL', 'Detail', null, true, 1, 'SCREEN', null, null, null, '2022-06-09 16:24:20.328077 +00:00', '2022-06-09 16:24:20.328077 +00:00', null, null, null),
        (4, 'MENU_DELETE', 'Delete', null, true, 1, 'ACTION', null, null, null, '2022-06-09 16:24:20.328077 +00:00', '2022-06-09 16:24:20.328077 +00:00', null, null, null),
        (5, 'MENU_UPDATE', 'Save', null, true, 3, 'ACTION', null, null, null, '2022-06-09 16:24:20.328077 +00:00', '2022-06-09 16:24:20.328077 +00:00', null, null, null),
        (6, 'MENU_CREATE', 'Save', null, true, 2, 'ACTION', null, null, null, '2022-06-09 16:27:29.237108 +00:00', '2022-06-09 16:27:29.237108 +00:00', null, null, null),
        (7, 'ROLE_MAIN', 'Role', null, true, null, 'SCREEN', null, null, null, '2022-06-09 16:30:44.373225 +00:00', '2022-06-09 16:30:44.373225 +00:00', null, null, null),
        (8, 'ROLE_ADD', 'Add', null, true, 7, 'ACTION', null, null, null, '2022-06-09 16:24:20.328077 +00:00', '2022-06-09 16:24:20.328077 +00:00', null, null, null),
        (9, 'ROLE_DETAIL', 'Detail', null, true, 7, 'SCREEN', null, null, null, '2022-06-09 16:24:20.328077 +00:00', '2022-06-09 16:24:20.328077 +00:00', null, null, null),
        (10, 'ROLE_DELETE', 'Delete', null, true, 7, 'ACTION', null, null, null, '2022-06-09 16:24:20.328077 +00:00', '2022-06-09 16:24:20.328077 +00:00', null, null, null),
        (11, 'ROLE_UPDATE', 'Save', null, true, 9, 'ACTION', null, null, null, '2022-06-09 16:24:20.328077 +00:00', '2022-06-09 16:24:20.328077 +00:00', null, null, null),
        (12, 'ROLE_CREATE', 'Save', null, true, 8, 'ACTION', null, null, null, '2022-06-09 16:42:48.103682 +00:00', '2022-06-09 16:42:48.103682 +00:00', null, null, null),
        (13, 'USER_MAIN', 'User', null, true, null, 'SCREEN', null, null, null, '2022-06-09 16:45:41.430584 +00:00', '2022-06-09 16:45:41.430584 +00:00', null, null, null),
        (14, 'USER_ADD', 'Add', null, true, 13, 'ACTION', null, null, null, '2022-06-09 16:45:41.430584 +00:00', '2022-06-09 16:45:41.430584 +00:00', null, null, null),
        (15, 'USER_DETAIL', 'Detail', null, true, 13, 'SCREEN', null, null, null, '2022-06-09 16:45:41.430584 +00:00', '2022-06-09 16:45:41.430584 +00:00', null, null, null),
        (16, 'USER_DELETE', 'Delete', null, true, 13, 'ACTION', null, null, null, '2022-06-09 16:45:41.430584 +00:00', '2022-06-09 16:45:41.430584 +00:00', null, null, null),
        (17, 'USER_UPDATE', 'Save', null, true, 15, 'ACTION', null, null, null, '2022-06-09 16:45:41.430584 +00:00', '2022-06-09 16:45:41.430584 +00:00', null, null, null),
        (18, 'USER_CREATE', 'Save', null, true, 14, 'ACTION', null, null, null, '2022-06-09 16:45:41.430584 +00:00', '2022-06-09 16:45:41.430584 +00:00', null, null, null),
        (19, 'USER_ROLE_ASSIGN', 'Role Assign', null, true, 15, 'SCREEN', null, null, null, '2022-06-09 16:45:41.430584 +00:00', '2022-06-09 16:45:41.430584 +00:00', null, null, null),
        (20, 'ROLE_MENU_ASSIGN', 'Menu Assign', null, true, 9, 'SCREEN', null, null, null, '2022-06-09 16:45:41.430584 +00:00', '2022-06-09 16:45:41.430584 +00:00', null, null, null);`

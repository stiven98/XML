


INSERT INTO roles (id, role) VALUES ('6c32abef-1f1a-4134-ba63-44c1d9e3ff7e', 'ROLE_ADMIN');
INSERT INTO roles (id, role) VALUES ('93baa4ef-95cb-4ff0-bd62-f5f26d6a0c69', 'ROLE_SYSTEM_USER');
INSERT INTO roles (id, role) VALUES ('437e1dd2-9cc2-4c87-81ba-0f1b52978bd6', 'ROLE_AGENT');

INSERT INTO privileges (privilege_id, name, role_id) VALUES ('b2cc1183-5af6-4a2c-932b-28559a5fc7df', 'READ_ACCESS', '6c32abef-1f1a-4134-ba63-44c1d9e3ff7e');
INSERT INTO privileges (privilege_id, name, role_id) VALUES ('c828aa64-6862-4b70-9803-b6e4b51c3db7', 'WRITE_ACCESS', '6c32abef-1f1a-4134-ba63-44c1d9e3ff7e');
INSERT INTO privileges (privilege_id, name, role_id) VALUES ('28bc8ec8-a642-430d-89c0-097afe7a24cb', 'READ_ACCESS', '93baa4ef-95cb-4ff0-bd62-f5f26d6a0c69');
INSERT INTO privileges (privilege_id, name, role_id) VALUES ('35cbedd2-fc09-45a7-8283-7bde4316d76c', 'WRITE_ACCESS', '93baa4ef-95cb-4ff0-bd62-f5f26d6a0c69');
INSERT INTO privileges (privilege_id, name, role_id) VALUES ('41ac8298-5619-4a35-bb66-d21cd5b8e6b0', 'READ_ACCESS', '437e1dd2-9cc2-4c87-81ba-0f1b52978bd6');
INSERT INTO privileges (privilege_id, name, role_id) VALUES ('32740915-4ea5-4f87-8fec-f046572f34a1', 'WRITE_ACCESS', '437e1dd2-9cc2-4c87-81ba-0f1b52978bd6');


INSERT INTO authentication_data (id, active, password, username) VALUES ('69b0597e-4a63-49e5-ae40-5b159ada82b9', true, '$2y$10$szTo3OrMpAUX0kIvWHh0seRntNn/GG6zBWIRnK.DJ7y.zItJRLYO2', 'acasimic');
INSERT INTO authentication_data (id, active, password, username) VALUES ('965208b9-287b-4da5-b772-73df5e74ebbc', true, '$2y$10$szTo3OrMpAUX0kIvWHh0seRntNn/GG6zBWIRnK.DJ7y.zItJRLYO2', 'jovanbosnic');
INSERT INTO authentication_data (id, active, password, username) VALUES ('4579daae-1567-42d5-a25c-1a3818077c84', true, '$2y$10$szTo3OrMpAUX0kIvWHh0seRntNn/GG6zBWIRnK.DJ7y.zItJRLYO2', 'djordjije');
INSERT INTO authentication_data (id, active, password, username) VALUES ('5cb65bc8-6130-4436-a1f9-ad4778f112bc', true, '$2y$10$szTo3OrMpAUX0kIvWHh0seRntNn/GG6zBWIRnK.DJ7y.zItJRLYO2', 'aleksandar');
INSERT INTO authentication_data (id, active, password, username) VALUES ('708b65de-fb77-4934-bfd0-d14161a74905', true, '$2y$10$szTo3OrMpAUX0kIvWHh0seRntNn/GG6zBWIRnK.DJ7y.zItJRLYO2', 'marko');
INSERT INTO authentication_data (id, active, password, username) VALUES ('0cf8a7ff-7bb5-48f0-a834-7b07eb306f90', true, '$2y$10$szTo3OrMpAUX0kIvWHh0seRntNn/GG6zBWIRnK.DJ7y.zItJRLYO2', 'janko_98');
INSERT INTO authentication_data (id, active, password, username) VALUES ('be71d1da-0749-480f-a563-dcc35a14e542', true, '$2y$10$szTo3OrMpAUX0kIvWHh0seRntNn/GG6zBWIRnK.DJ7y.zItJRLYO2', 'deki_99');
INSERT INTO authentication_data (id, active, password, username) VALUES ('d3ea863d-350e-44f2-bd6e-809aa7100476', true, '$2y$10$szTo3OrMpAUX0kIvWHh0seRntNn/GG6zBWIRnK.DJ7y.zItJRLYO2', 'milica00');






INSERT INTO user_role (user_id, role_id) VALUES ('69b0597e-4a63-49e5-ae40-5b159ada82b9', '6c32abef-1f1a-4134-ba63-44c1d9e3ff7e');
INSERT INTO user_role (user_id, role_id) VALUES ('965208b9-287b-4da5-b772-73df5e74ebbc', '93baa4ef-95cb-4ff0-bd62-f5f26d6a0c69');
INSERT INTO user_role (user_id, role_id) VALUES ('4579daae-1567-42d5-a25c-1a3818077c84', '93baa4ef-95cb-4ff0-bd62-f5f26d6a0c69');
INSERT INTO user_role (user_id, role_id) VALUES ('5cb65bc8-6130-4436-a1f9-ad4778f112bc', '93baa4ef-95cb-4ff0-bd62-f5f26d6a0c69');
INSERT INTO user_role (user_id, role_id) VALUES ('708b65de-fb77-4934-bfd0-d14161a74905', '93baa4ef-95cb-4ff0-bd62-f5f26d6a0c69');
INSERT INTO user_role (user_id, role_id) VALUES ('0cf8a7ff-7bb5-48f0-a834-7b07eb306f90', '93baa4ef-95cb-4ff0-bd62-f5f26d6a0c69');
INSERT INTO user_role (user_id, role_id) VALUES ('be71d1da-0749-480f-a563-dcc35a14e542', '93baa4ef-95cb-4ff0-bd62-f5f26d6a0c69');
INSERT INTO user_role (user_id, role_id) VALUES ('d3ea863d-350e-44f2-bd6e-809aa7100476', '93baa4ef-95cb-4ff0-bd62-f5f26d6a0c69');



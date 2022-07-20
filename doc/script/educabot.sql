PGDMP         #                z            educabot #   12.9 (Ubuntu 12.9-0ubuntu0.20.04.1) #   12.9 (Ubuntu 12.9-0ubuntu0.20.04.1) -    �           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            �           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            �           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            �           1262    17851    educabot    DATABASE     z   CREATE DATABASE educabot WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'es_AR.UTF-8' LC_CTYPE = 'es_AR.UTF-8';
    DROP DATABASE educabot;
                postgres    false            �            1259    18785    driver_trip    TABLE     `   CREATE TABLE public.driver_trip (
    trip_id bigint NOT NULL,
    driver_id bigint NOT NULL
);
    DROP TABLE public.driver_trip;
       public         heap    postgres    false            �            1259    18728    drivers    TABLE     �   CREATE TABLE public.drivers (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    driver_file bigint,
    description text,
    user_id bigint NOT NULL
);
    DROP TABLE public.drivers;
       public         heap    postgres    false            �            1259    18726    drivers_id_seq    SEQUENCE     w   CREATE SEQUENCE public.drivers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 %   DROP SEQUENCE public.drivers_id_seq;
       public          postgres    false    207            �           0    0    drivers_id_seq    SEQUENCE OWNED BY     A   ALTER SEQUENCE public.drivers_id_seq OWNED BY public.drivers.id;
          public          postgres    false    206            �            1259    17900    roles    TABLE     �   CREATE TABLE public.roles (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name character varying(255)
);
    DROP TABLE public.roles;
       public         heap    postgres    false            �            1259    17898    roles_id_seq    SEQUENCE     u   CREATE SEQUENCE public.roles_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.roles_id_seq;
       public          postgres    false    203            �           0    0    roles_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.roles_id_seq OWNED BY public.roles.id;
          public          postgres    false    202            �            1259    18746    trips    TABLE     �   CREATE TABLE public.trips (
    id bigint NOT NULL,
    start_time timestamp with time zone NOT NULL,
    end_time timestamp with time zone,
    driver_id bigint,
    finished boolean DEFAULT false
);
    DROP TABLE public.trips;
       public         heap    postgres    false            �            1259    18744    trips_id_seq    SEQUENCE     u   CREATE SEQUENCE public.trips_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.trips_id_seq;
       public          postgres    false    209            �           0    0    trips_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.trips_id_seq OWNED BY public.trips.id;
          public          postgres    false    208            �            1259    18711    users    TABLE     �   CREATE TABLE public.users (
    id bigint NOT NULL,
    name character varying(255),
    email character varying(255),
    password text NOT NULL,
    role_id bigint NOT NULL
);
    DROP TABLE public.users;
       public         heap    postgres    false            �            1259    18709    users_id_seq    SEQUENCE     u   CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.users_id_seq;
       public          postgres    false    205            �           0    0    users_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;
          public          postgres    false    204            R           2604    18731 
   drivers id    DEFAULT     h   ALTER TABLE ONLY public.drivers ALTER COLUMN id SET DEFAULT nextval('public.drivers_id_seq'::regclass);
 9   ALTER TABLE public.drivers ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    206    207    207            P           2604    17903    roles id    DEFAULT     d   ALTER TABLE ONLY public.roles ALTER COLUMN id SET DEFAULT nextval('public.roles_id_seq'::regclass);
 7   ALTER TABLE public.roles ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    203    202    203            S           2604    18749    trips id    DEFAULT     d   ALTER TABLE ONLY public.trips ALTER COLUMN id SET DEFAULT nextval('public.trips_id_seq'::regclass);
 7   ALTER TABLE public.trips ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    208    209    209            Q           2604    18714    users id    DEFAULT     d   ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);
 7   ALTER TABLE public.users ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    204    205    205            �          0    18785    driver_trip 
   TABLE DATA           9   COPY public.driver_trip (trip_id, driver_id) FROM stdin;
    public          postgres    false    210   �0       �          0    18728    drivers 
   TABLE DATA           l   COPY public.drivers (id, created_at, updated_at, deleted_at, driver_file, description, user_id) FROM stdin;
    public          postgres    false    207   �0       �          0    17900    roles 
   TABLE DATA           M   COPY public.roles (id, created_at, updated_at, deleted_at, name) FROM stdin;
    public          postgres    false    203   Q1       �          0    18746    trips 
   TABLE DATA           N   COPY public.trips (id, start_time, end_time, driver_id, finished) FROM stdin;
    public          postgres    false    209   �1       �          0    18711    users 
   TABLE DATA           C   COPY public.users (id, name, email, password, role_id) FROM stdin;
    public          postgres    false    205   �2       �           0    0    drivers_id_seq    SEQUENCE SET     <   SELECT pg_catalog.setval('public.drivers_id_seq', 6, true);
          public          postgres    false    206            �           0    0    roles_id_seq    SEQUENCE SET     :   SELECT pg_catalog.setval('public.roles_id_seq', 2, true);
          public          postgres    false    202            �           0    0    trips_id_seq    SEQUENCE SET     ;   SELECT pg_catalog.setval('public.trips_id_seq', 14, true);
          public          postgres    false    208            �           0    0    users_id_seq    SEQUENCE SET     :   SELECT pg_catalog.setval('public.users_id_seq', 7, true);
          public          postgres    false    204            c           2606    18789    driver_trip driver_trip_pkey 
   CONSTRAINT     j   ALTER TABLE ONLY public.driver_trip
    ADD CONSTRAINT driver_trip_pkey PRIMARY KEY (trip_id, driver_id);
 F   ALTER TABLE ONLY public.driver_trip DROP CONSTRAINT driver_trip_pkey;
       public            postgres    false    210    210            ]           2606    18736    drivers drivers_pkey 
   CONSTRAINT     R   ALTER TABLE ONLY public.drivers
    ADD CONSTRAINT drivers_pkey PRIMARY KEY (id);
 >   ALTER TABLE ONLY public.drivers DROP CONSTRAINT drivers_pkey;
       public            postgres    false    207            X           2606    17905    roles roles_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.roles DROP CONSTRAINT roles_pkey;
       public            postgres    false    203            a           2606    18751    trips trips_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.trips
    ADD CONSTRAINT trips_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.trips DROP CONSTRAINT trips_pkey;
       public            postgres    false    209            [           2606    18719    users users_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pkey;
       public            postgres    false    205            ^           1259    18743    idx_drivers_deleted_at    INDEX     P   CREATE INDEX idx_drivers_deleted_at ON public.drivers USING btree (deleted_at);
 *   DROP INDEX public.idx_drivers_deleted_at;
       public            postgres    false    207            _           1259    27690    idx_drivers_driver_file    INDEX     Y   CREATE UNIQUE INDEX idx_drivers_driver_file ON public.drivers USING btree (driver_file);
 +   DROP INDEX public.idx_drivers_driver_file;
       public            postgres    false    207            U           1259    17907    idx_roles_deleted_at    INDEX     L   CREATE INDEX idx_roles_deleted_at ON public.roles USING btree (deleted_at);
 (   DROP INDEX public.idx_roles_deleted_at;
       public            postgres    false    203            V           1259    17906    idx_roles_name    INDEX     G   CREATE UNIQUE INDEX idx_roles_name ON public.roles USING btree (name);
 "   DROP INDEX public.idx_roles_name;
       public            postgres    false    203            Y           1259    18725    idx_users_email    INDEX     I   CREATE UNIQUE INDEX idx_users_email ON public.users USING btree (email);
 #   DROP INDEX public.idx_users_email;
       public            postgres    false    205            h           2606    18795 !   driver_trip fk_driver_trip_driver    FK CONSTRAINT     �   ALTER TABLE ONLY public.driver_trip
    ADD CONSTRAINT fk_driver_trip_driver FOREIGN KEY (driver_id) REFERENCES public.drivers(id);
 K   ALTER TABLE ONLY public.driver_trip DROP CONSTRAINT fk_driver_trip_driver;
       public          postgres    false    210    207    2909            g           2606    18790    driver_trip fk_driver_trip_trip    FK CONSTRAINT     ~   ALTER TABLE ONLY public.driver_trip
    ADD CONSTRAINT fk_driver_trip_trip FOREIGN KEY (trip_id) REFERENCES public.trips(id);
 I   ALTER TABLE ONLY public.driver_trip DROP CONSTRAINT fk_driver_trip_trip;
       public          postgres    false    210    209    2913            f           2606    27696    trips fk_drivers_trip    FK CONSTRAINT     x   ALTER TABLE ONLY public.trips
    ADD CONSTRAINT fk_drivers_trip FOREIGN KEY (driver_id) REFERENCES public.drivers(id);
 ?   ALTER TABLE ONLY public.trips DROP CONSTRAINT fk_drivers_trip;
       public          postgres    false    207    209    2909            d           2606    27685    users fk_roles_users    FK CONSTRAINT     s   ALTER TABLE ONLY public.users
    ADD CONSTRAINT fk_roles_users FOREIGN KEY (role_id) REFERENCES public.roles(id);
 >   ALTER TABLE ONLY public.users DROP CONSTRAINT fk_roles_users;
       public          postgres    false    2904    203    205            e           2606    27691    drivers fk_users_driver    FK CONSTRAINT     v   ALTER TABLE ONLY public.drivers
    ADD CONSTRAINT fk_users_driver FOREIGN KEY (user_id) REFERENCES public.users(id);
 A   ALTER TABLE ONLY public.drivers DROP CONSTRAINT fk_users_driver;
       public          postgres    false    207    205    2907            �      x������ � �      �   b   x���1�0D��{
_���,���Pb$*$�8�4N���?*��<�UKU����ڟ��hӼ��O����ݮ;P8��j%���~aa�5����'      �   "   x�3��Ĕ��<.#�@JQfYjW� �<	�      �   C  x����m�PDϡ
�����6rN�_1 �X�\���3��F�}��	˲!���rQ�ԋ�U�XU����&8�)?	f��-���=^���B�"�d��`J� �&<Rl�
�aD'a�=ZĠ>p�a�),B����c���a�i����n��&8tӑ�0ֈL�GlB^<�++�@�G��R�H�=�J�����AX))6�	x���9�LB���AzA��:I��C�!����7����B�T�8� zN��7���a �؋ڂ�M�4�/7�݄��/B3�$}0�z6]k̟_=��Dz�tXh��/�?5��i�]�      �   �   x�3�LKMI�,I-.qH�M���K���T1JT10Q�--/Jq�-��v�v�4L-*��LL�70�1�2�p6	7�����(tJ�)�4�2��/.�M��b�T�k�e3V�75�� �N̥��1z\\\ #�U�     
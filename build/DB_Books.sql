PGDMP         +    	            {            DB_Books    15.2    15.2     �           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            �           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            �           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            �           1262    16398    DB_Books    DATABASE     �   CREATE DATABASE "DB_Books" WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'English_United States.1252';
    DROP DATABASE "DB_Books";
                postgres    false            �            1259    16422    books    TABLE       CREATE TABLE public.books (
    id integer NOT NULL,
    title character varying(200),
    author character varying(200),
    des character varying(200),
    create_at timestamp without time zone,
    update_at timestamp without time zone,
    delete_at timestamp without time zone
);
    DROP TABLE public.books;
       public         heap    postgres    false            �            1259    16421    books_id_seq    SEQUENCE     �   CREATE SEQUENCE public.books_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.books_id_seq;
       public          postgres    false    215            �           0    0    books_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.books_id_seq OWNED BY public.books.id;
          public          postgres    false    214            e           2604    16425    books id    DEFAULT     d   ALTER TABLE ONLY public.books ALTER COLUMN id SET DEFAULT nextval('public.books_id_seq'::regclass);
 7   ALTER TABLE public.books ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    214    215    215            �          0    16422    books 
   TABLE DATA           X   COPY public.books (id, title, author, des, create_at, update_at, delete_at) FROM stdin;
    public          postgres    false    215          �           0    0    books_id_seq    SEQUENCE SET     ;   SELECT pg_catalog.setval('public.books_id_seq', 16, true);
          public          postgres    false    214            g           2606    16429    books books_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.books
    ADD CONSTRAINT books_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.books DROP CONSTRAINT books_pkey;
       public            postgres    false    215            �   v   x�34��IM,�Sp��V�R(1��-�s�̫�p@���E��ىy
I��ŉ
��E��E��@��Լt ���W�Z��� �gd`d�k`�kh�`haelied�gnjjll��D\1z\\\ HL#     
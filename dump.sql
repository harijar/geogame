--
-- PostgreSQL database dump
--

-- Dumped from database version 16.0 (Debian 16.0-1.pgdg120+1)
-- Dumped by pg_dump version 16.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: countries; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.countries (
    id integer NOT NULL,
    name text,
    aliases text[],
    un_not_member text,
    unrecognised text,
    capital text,
    religion text,
    religion_perc text,
    population integer,
    area double precision,
    gdp integer,
    gdp_per_capita integer,
    hdi double precision,
    independent_from text,
    agricultural_sector double precision,
    industrial_sector double precision,
    service_sector double precision,
    northernmost double precision,
    southernmost double precision,
    easternmost double precision,
    westernmost double precision,
    hemisphere_lat integer,
    hemisphere_long integer,
    monarchy boolean,
    landlocked boolean,
    island boolean
);


ALTER TABLE public.countries OWNER TO postgres;

--
-- Name: ethnic_groups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.ethnic_groups (
    country_id integer,
    name text,
    percentage double precision
);


ALTER TABLE public.ethnic_groups OWNER TO postgres;

--
-- Name: funfacts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.funfacts (
    country_id integer,
    text text
);


ALTER TABLE public.funfacts OWNER TO postgres;

--
-- Name: languages; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.languages (
    country_id integer,
    name text
);


ALTER TABLE public.languages OWNER TO postgres;

--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO postgres;

--
-- Data for Name: countries; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.countries (id, name, aliases, un_not_member, unrecognised, capital, religion, religion_perc, population, area, gdp, gdp_per_capita, hdi, independent_from, agricultural_sector, industrial_sector, service_sector, northernmost, southernmost, easternmost, westernmost, hemisphere_lat, hemisphere_long, monarchy, landlocked, island) FROM stdin;
1	Sudan	{sudan}			Khartoum	 Sunni Islam	97	41984500	1731671	35867	786	0.508	United Kingdom	39.6	2.6	57.8	22.12	9.3	38.35	21.5	0	0	f	f	f
2	Comoros	{comoros,"comoro islands"}			Moroni		0	758316	1862	1340	1631	0.558	France	47.7	11.8	40.5	-11.22	-12.25	44.32	43.12	1	0	f	f	t
3	Malawi	{malawi}			Lilongwe	Christianity	82.3	21507723	94080	12199	613	0.512	United Kingdom	28.6	15.4	56	-9.22	-17.07	35.55	32.4	1	0	f	t	f
4	Philippines	{philippines}			Manila	Christianity	90.1	110750000	298170	394086	3461	0.699	Spanish Empire	9.6	30.6	59.8	21.07	4.35	126.34	114.17	0	0	f	f	t
5	Poland	{poland}			Warsaw	Christianity	72.2	37691000	311888	679442	17736	0.876	Russian Republic	2.4	40.2	57.4	54.5	49	24.09	14.07	0	0	f	f	f
6	Romania	{romania}			Bucharest	Christianity	84.8	19051562	231291	284086	14698	0.821	Ottoman Empire	4.2	33.2	62.6	48.15	43.38	29.4	20.19	0	0	f	f	f
7	Kenya	{kenya}			Nairobi	Christianity	85.5	51526000	569140	110347	2082	0.575	United Kingdom	34.5	17.8	47.5	5.35	-4.43	41	35	2	0	f	f	f
8	Mexico	{mexico}			Mexico City	Roman Catholic	56	129202482	1943945	1272839	10046	0.758	Spanish Empire	3.6	31.9	64.5	32.43	14.32	-86.42	-117.02	0	1	f	f	f
9	Eswatini	{eswatini}			style	Christianity	89.3	1223362	17204	4732	3969	0.597	United Kingdom	6.5	45	48.6	-25.44	-27.19	32.08	30.48	1	0	t	t	f
10	Iran	{iran}			Tehran	Islam	0	85365900	1531595	594892	6766	0.774	Imperial State of Iran	9.6	35.3	55	39.47	25.04	63.18	44.02	0	0	f	f	f
11	Panama	{panama}			Panama City	Christianity	91.5	4337406	74340	63605	14618	0.805	Spanish Empire	2.4	15.7	82	9.38	7.11	-77.09	-83.04	0	1	f	f	f
12	Cape Verde	{"cape verde"}			Praia	Christianity	81.7	491233	4033	1936	3293	0.662	Portugal	8.9	17.5	73.7	17.12	14.48	-22.4	-25.25	0	1	f	f	t
13	Czech Republic	{"czech republic",czechia}			Prague	no religion	56.9	10873553	77171	281778	26809	0.889	Austria-Hungary	2.3	36.9	60.8	51.02	48.33	18.51	12.05	0	0	f	t	f
14	Mauritania	{mauritania}			Nouakchott	Islam	0	4475683	1025520	9996	2166	0.556	France	27.8	29.3	42.9	27.19	14.43	-4.45	-17.03	0	1	f	f	f
15	Austria	{austria}			Vienna	Christianity	64.1	9129652	82445	480368	53840	0.916		1.3	28.4	70.3	49.01	46.23	17.09	9.31	0	0	f	t	f
16	The Bahamas	{"the bahamas",bahamas}			Nassau	Christianity	93	397360	10010	0	27478	0.812	United Kingdom	2.3	7.7	90	27.16	20.55	-72.45	-79.15	0	1	t	f	f
17	Bosnia and Herzegovina	{"bosnia and herzegovina",bosnia}			Sarajevo	Islam	51	3277082	51187	23365	7143	0.78	Yugoslavia	6.8	28.9	64.3	45.16	42.34	19.3	15.5	0	0	f	f	f
18	Brazil	{brazil}			Brasília	Christianity	88.8	203062512	8460415	1608981	7507	0.754	Portugal	6.6	20.7	72.7	5.15	-33.45	-34.47	-73.59	2	1	f	f	f
19	Burkina Faso	{"burkina faso"}			Ouagadougou		0	22185654	273602	19738	893	0.449	France	31	23.9	44.9	15.05	9.24	2.25	-5.3	0	2	f	t	f
20	Vietnam	{vietnam}			Hanoi	no religion	73.7	100000000	310070	366138	3756	0.703	Empire of Japan	15.3	33.3	51.3	23.27	8.35	109.3	102.08	0	0	f	f	f
21	Serbia	{serbia}			Belgrade	Christianity	86.6	6641197	88246	63068	8643	0.803	Ottoman Empire	9.8	41.1	49.1	46.11	42.13	23.01	18.51	0	0	f	t	f
22	Cuba	{cuba}			Havana	Christianity	58.9	11089511	109884	126694	11255	0.764	Spanish Empire	4	22.7	73.4	23.12	19.5	-74.08	-84.58	0	1	f	f	t
23	Guinea	{guinea}			Conakry		0	13261638	245717	16036	1185	0.465	France	19.8	32.1	48.1	12.42	7.12	-7.45	-15.03	0	1	f	f	f
24	Israel	{israel}		Partially unrecognised	Jerusalem	Judaism	0	9802400	20330	481591	54111	0.919	United Kingdom	2.4	26.5	69.5	33.17	29.3	35.4	34.15	0	0	f	f	f
25	Mozambique	{mozambique}			Maputo	Christianity	55.8	32419747	786380	15777	492	0.446	Portugal	23.9	19.3	56.8	-10.27	-26.52	40.45	30.15	1	0	f	f	f
26	Papua New Guinea	{"papua new guinea"}			Port Moresby	Christianity	95.5	11781559	452860	26595	2673	0.558	Australia	22.1	42.9	35	-2.35	-10.42	150.52	140.55	1	0	t	f	t
27	Brunei	{brunei}			Bandar Seri Begawan	Sunni Islam	80.9	445400	5265	14006	31449	0.829	United Kingdom	1.2	56.6	42.3	5.04	4	115.25	114.05	0	0	t	f	f
28	Slovakia	{slovakia}			Bratislava	Christianity	68.8	5426352	48105	116527	21390	0.848	Czechoslovakia	3.8	35	61.2	49.36	47.44	22.3	16.5	0	0	f	t	f
29	Burundi	{burundi}			Gitega	link	93.4	12837740	25680	3900	311	0.426	Belgium	39.5	16.4	44.2	-2.18	-4.28	30.5	28.59	1	0	f	t	f
30	Ghana	{ghana}			Accra	Christianity	71.3	30832019	228567	79083	2409	0.632	United Kingdom	18.3	24.5	57.2	11.1	4.44	1.1	-3.15	0	2	f	f	f
31	Iceland	{iceland}			Reykjavík	Christianity	72.4	396960	100250	25602	69133	0.959	Denmark	5.8	19.7	74.6	66.32	63.23	-13.14	-24.32	0	1	f	f	t
32	Uganda	{uganda}			Kampala	Christianity	84.4	42885900	197100	42661	930	0.525	United Kingdom	28.2	21.1	50.7	4.14	-1.29	35.02	29.34	2	0	f	t	f
33	Zambia	{zambia}			Lusaka	Christianity	0	19610769	743398	21313	1095	0.565	United Kingdom	7.5	35.3	57	-8.12	-18.05	32.5	22	1	0	f	t	f
34	Uzbekistan	{uzbekistan}			Tashkent	Islam	96.1	36599764	425400	69239	2032	0.727	Soviet Union	17.9	33.7	48.5	45.36	37.11	73.2	56	0	0	f	t	f
35	Benin	{benin}			Porto-Novo	Christianity	52.2	12606998	114305	17688	1361	0.525	France	26.1	22.8	51.1	12.25	6.14	3.5	0.5	0	0	f	f	f
36	France	{france}			Paris	Christianity	50	68184000	640427	2957880	44229	0.903		1.7	20.2	79.2	51.05	42.2	8.13	-4.47	0	2	f	f	f
73	Mali	{mali}			Bamako	Islam	95	22395489	1220190	19157	875	0.428	France	41.8	18.1	40.5	25.04	10.09	4	-12	0	2	f	t	f
37	Kyrgyzstan	{kyrgyzstan}			Bishkek	Islam	90	7100000	191801	8741	1339	0.692	Soviet Union	14.6	31.2	54.2	43.16	39.1	80.2	69.2	0	0	f	t	f
38	Federated States of Micronesia	{"federated states of micronesia"}			Palikir	Christianity	95.3	105754	702	0	3573	0.628	United States	26.3	18.9	54.8	10.05	1.01	163.01	137.3	0	0	f	f	f
39	Seychelles	{seychelles}			Victoria	Christianity	94.7	100447	455	1287	12085	0.785	United Kingdom	2.5	13.8	83.7	-3.43	-10.13	56.16	46.29	1	0	f	f	t
40	Jamaica	{jamaica}			Kingston	Christianity	68.9	2825544	10831	14658	5184	0.709	United Kingdom	7	21.1	71.9	18.31	17.42	-76.12	-78.22	0	1	t	f	t
41	Moldova	{moldova}			Chișinău	Christianity	91.8	2512758	32891	13680	4468	0.767	Soviet Union	17.7	20.3	62	48.28	45.28	30.1	26.4	0	0	f	t	f
42	Senegal	{senegal}			Dakar		0	18275743	192530	27625	1637	0.511	France	16.9	24.3	58.8	16.41	12.18	-11.22	-17.27	0	1	f	f	f
43	Turkey	{turkey}			Ankara	Islam	0	85279553	769632	819034	9661	0.838		6.8	32.3	60.7	42.06	35.49	44.49	26.03	0	0	f	f	f
44	Vanuatu	{vanuatu}			Port Vila	Christianity	93.4	301295	12189	981	3073	0.607	Papua New Guinea	27.3	11.8	60.8	-13.07	-20.15	170.13	166.33	1	0	f	f	t
45	Croatia	{croatia}			Zagreb	Christianity	87.4	3855641	55974	68955	16983	0.858	Yugoslavia	3.7	26.2	70.1	46.33	42.23	19.27	13.3	0	0	f	f	f
46	Denmark	{denmark}			Copenhagen	Christianity	75.8	5959464	42434	398303	68037	0.948		1.3	22.9	75.8	57.45	54.34	15.11	8.05	0	0	t	f	f
47	Saudi Arabia	{"saudi arabia"}			Riyadh	Islam	93	32175224	2149690	833541	23186	0.875		2.6	44.2	53.2	32.14	16.23	55.3	34.5	0	0	t	f	f
48	Uruguay	{uruguay}			Montevideo	Christianity	54.3	3566550	175015	59318	17313	0.809	Empire of Brazil	6.2	24.1	69.7	-30.05	-34.58	-53.5	-58.25	1	1	f	f	f
49	Malaysia	{malaysia}			Kuala Lumpur	Sunni Islam	63.5	33379500	329613	372702	11101	0.803	United Kingdom	8.8	37.6	53.6	7.15	0.51	119.1	99.37	0	0	t	f	f
50	Namibia	{namibia}			Windhoek	Christianity	87.9	2641857	823290	12236	4836	0.615	South Africa	6.7	26.3	67	-16.56	-28.58	25.15	11.45	1	0	f	f	f
51	Singapore	{singapore}			Singapore	Buddhism	31.1	5917600	716	396992	66822	0.939	Malaysia	0	24.8	75.2	1.28	1.09	104.24	103.36	0	0	f	f	t
52	Argentina	{argentina}			Buenos Aires	Christianity	58.9	46654581	2736690	487227	10761	0.842	Spanish Empire	10.8	28.1	61.1	-21.48	-52.24	-53.41	-73.3	1	1	f	f	f
53	China	{china}			Beijing	no religion	74.5	1411750000	9326410	17734131	12437	0.768	Republic of China	7.9	40.5	51.6	53.34	20.14	134.46	73.33	0	0	f	f	f
54	Guyana	{guyana}			Georgetown	Christianity	62.7	743699	196849	8044	9999	0.714	Portugal	15.4	15.3	69.3	8.24	1.11	-56.3	-61.25	0	1	f	f	f
55	Laos	{laos}			Vientiane	Buddhism	66	7443000	230800	19074	2569	0.607		20.9	33.2	45.9	22.29	13.55	106.4	100.08	0	0	f	t	f
56	Liberia	{liberia}			Monrovia	Christianity	85.1	5248621	96320	2445	471	0.481	American Colonization Society	34	13.8	52.2	8.33	4.21	-7.22	-11.3	0	1	f	f	f
57	Somalia	{somalia}			Mogadishu	Islam	0	18143379	627337	7628	447	0	UN trusteeship	60.2	7.4	32.5	11.59	-1.4	51.27	41	2	0	f	f	f
58	Zimbabwe	{zimbabwe}			Harare	Christianity	84.1	15178979	386847	24118	1508	0.593	United Kingdom	12	22.2	65.8	-15.37	-22.25	33	25.2	1	0	f	f	f
59	Qatar	{qatar}			Doha	Islam	65.5	2656032	11586	179571	66799	0.855	Ottoman Empire	0.2	50.3	49.5	26.09	24.28	51.39	50.45	0	0	t	f	f
60	Yemen	{yemen}			Sanaa	Islam	99.1	31888698	527968	9947	302	0.455	United Kingdom	20.3	11.8	67.9	18.59	12.36	53.1	41.49	0	0	f	f	f
61	Armenia	{armenia}			Yerevan	Christianity	0	2981200	28342	13861	4967	0.759	Transcaucasian Democratic Federative Republic	16.7	28.2	54.8	41.17	38.5	46.37	43.27	0	0	f	t	f
62	Central African Republic	{"central african republic",car}			Bangui	Christianity	73.2	5633412	622984	2518	461	0.404	France	43.2	16	40.8	11.01	2.13	27.28	14.2	0	0	f	t	f
63	Costa Rica	{"costa rica"}			San José	Christianity	72.6	5262225	51060	64282	12472	0.809	Spanish Empire	5.5	20.6	73.9	11.13	8.04	-82.3	-85.5	0	1	f	f	f
64	Kazakhstan	{kazakhstan}			Astana	Islam	69.3	19944726	2699700	193018	10055	0.811	Soviet Union	4.7	34.1	61.2	55.43	40.34	87.3	50	0	0	f	t	f
65	Kuwait	{kuwait}			Kuwait City	Islam	76.7	4670713	17818	136642	32150	0.831	United Kingdom	0.4	58.7	40.9	30.06	28.32	48.4	46.34	0	0	t	f	f
66	Albania	{albania}			Tirana	Islam	59	2761785	27398	18260	6396	0.796	Ottoman Empire	21.7	24.2	54.1	42.3	39.39	21.4	19.16	0	0	f	f	f
67	Nepal	{nepal}			Kathmandu	Hinduism	81.19	29164578	143686	36207	1159	0.602		27	13.5	59.5	30.26	26.22	88.15	80.03	0	0	f	t	f
68	State of Palestine	{"state of palestine",palestine}	UN observer	Partially unrecognised		Islam	0	5483450	6000	0	3514	0.715	Israel	0	0	0	32.32	0	0	0	0	0	f	f	f
69	Sri Lanka	{"sri lanka"}			Sri Jayawardenepura Kotte	Buddhism	0	22037000	62732	85309	3918	0.782	United Kingdom	7.8	30.5	61.7	9.3	5.55	81.4	79.31	0	0	f	f	t
70	Belize	{belize}			Belmopan	Christianity	87.6	441471	22806	2492	6229	0.683	United Kingdom	10.3	21.6	68	18.27	15.53	-87.35	-89.12	0	1	t	f	f
71	Botswana	{botswana}			Gaborone	Christianity	79.1	2410338	566730	17615	6805	0.693	United Kingdom	1.8	27.5	70.6	-17.46	-26.54	29.2	20	1	0	f	t	f
72	Finland	{finland}			Helsinki	Christianity	68.6	5557272	303816	297302	53703	0.94	Soviet Union	2.7	28.2	69.1	70.05	59.48	31.35	20.32	0	0	f	f	f
74	Pakistan	{pakistan}			Islamabad	Islam	96.5	241499431	856690	342501	1480	0.544	United Kingdom	24.4	19.1	56.5	37.08	23.42	79.26	61	0	0	f	f	f
75	Hungary	{hungary}			Budapest	Christianity	42.5	9597085	89608	181848	18728	0.846	Hungarian People's Republic	3.9	31.3	64.8	48.35	45.45	22.55	16.8	0	0	f	t	f
76	Republic of Ireland	{"republic of ireland",ireland}			Dublin	Christianity	75.7	5281600	68883	0	101109	0.945		1.2	38.6	60.2	55.23	51.26	-5.59	-10.28	0	1	f	f	f
77	South Korea	{"south korea","republic of korea",rok}			Seoul	no religion	56.1	51439038	99909	1810966	34940	0.925	Empire of Japan	2.2	39.3	58.3	38.36	34.17	129.35	124.39	0	0	f	f	f
78	Andorra	{andorra}			Andorra la Vella	Christianity	0	83523	468	3325	42066	0.858		11.9	33.6	54.5	42.39	42.26	1.47	1.24	0	0	t	t	f
79	Bangladesh	{bangladesh}			Dhaka	Islam	91.04	169828911	134208	414907	2450	0.661	Pakistan	14.2	29.3	56.5	26.38	20.45	92.35	88.03	0	0	f	f	f
80	Dominican Republic	{"dominican republic",dominicana}			Santo Domingo	Christianity	66.7	10760028	48320	94243	8477	0.767	Spain	5.6	33	61.4	19.55	17.36	-68.2	-72	0	1	f	f	t
81	El Salvador	{"el salvador",salvador}			San Salvador	Christianity	84.1	6884888	20721	28737	4551	0.675	Spanish Empire	12	27.7	60.3	14.26	13.09	-87.42	-90.15	0	1	f	f	f
82	Gabon	{gabon}			Libreville	Christianity	82.2	2233272	257667	18521	7911	0.706	France	5	44.7	50.4	2.19	-3.59	14.3	8.42	2	0	f	f	f
83	Bhutan	{bhutan}			Thimphu	Buddhism	84.3	770276	38394	2381	3063	0.666		16.2	41.8	42	28.2	26.43	92.05	88.45	0	0	t	t	f
84	Colombia	{colombia}			Bogotá	Christianity	87	52215503	1038700	314464	6104	0.752	Spanish Empire	7.2	30.8	62.1	12.27	-4.14	-67	-79	2	1	f	f	f
85	Greece	{greece}			Athens	Christianity	93	10482487	130647	214874	20571	0.887	Ottoman Empire	4.1	16.9	79.1	41.43	36.23	26.38	19.59	0	0	f	f	f
86	Montenegro	{montenegro}			Podgorica	Christianity	76	616695	13452	5809	9252	0.832	Serbia and Montenegro	7.5	15.9	76.6	43.32	41.51	20.21	18.26	0	0	f	f	f
87	Samoa	{samoa}			Apia	Christianity	97.9	205557	2821	857	3919	0.707	New Zealand	10.4	23.6	66	-13.25	-14.04	-171.2	-172.48	1	1	f	f	t
88	Canada	{canada}			Ottawa	Christianity	0	40097761	9093507	1988336	52112	0.936	United Kingdom	1.6	28.2	70.2	72	41.57	-55.37	-141	0	1	t	f	f
89	Ecuador	{ecuador}			Quito	Christianity	86.7	16938986	256369	106166	5965	0.74	Spanish Empire	6.7	32.9	60.4	1.4	-5	-75.15	-80.05	2	1	f	f	f
90	Peru	{peru}			Lima	Christianity	94.5	33396698	1279996	223252	6622	0.762	Spanish Empire	7.6	32.7	59.9	-0.01	-18.21	-68.39	-81.19	1	1	f	f	f
91	Cameroon	{cameroon}			Yaoundé	Christianity	59.7	28088845	472710	45368	1668	0.576	United Kingdom	16.7	26.5	56.8	13.05	1.39	16.12	8.45	0	0	f	f	f
92	Jordan	{jordan}			Amman	Sunni Islam	95	11505800	88802	45244	4058	0.727	United Kingdom	4.5	28.8	66.6	33.22	29.11	39.2	34.55	0	0	t	f	f
93	Portugal	{portugal}			Lisbon	Christianity	84.8	10467366	91119	253663	24651	0.866	Spanish Empire	2.2	22.1	75.7	42.07	36.57	-6.11	-9.5	0	1	f	f	f
94	Tajikistan	{tajikistan}			Dushanbe	Islam	96.4	10077600	141510	8746	897	0.685	Soviet Union	28.6	25.5	45.9	41.02	36.4	75.05	67.25	0	0	f	t	f
95	Tuvalu	{tuvalu}			Funafuti	Christianity	94.8	10679	26	60	5370	0.641	United Kingdom	24.5	5.6	70	-5.39	-10.48	179.51	176.03	1	0	t	f	t
96	Oman	{oman}			Muscat	Islam	88.9	5113071	309500	88192	19509	0.816	Portugal	1.8	46.4	51.8	26.3	16.39	59.5	52	0	0	t	f	f
97	United States	{"united states",us,"the us",usa}			Washington, D.C.	Christianity	63	335623000	9147593	23315081	69185	0.921	Kingdom of Great Britain	0.9	19.1	80	49.23	25.07	-66.56	-124.43	0	1	f	f	f
98	Chad	{chad}			N'Djamena	Islam	55.1	17414717	1259200	16410	955	0.394	France	52.3	14.7	33.1	23.29	7.27	24	13.4	0	0	f	t	f
99	Iraq	{iraq}			Baghdad	Islam	95	43324000	437367	204004	4686	0.686	United Kingdom	3.3	51	45.8	37.23	29.04	48.32	38.5	0	0	f	f	f
100	Lebanon	{lebanon}			Beirut	Islam	0	5490000	10230	37945	6785	0.706	France	3.9	13.1	83	34.43	33.03	36.37	35.06	0	0	f	f	f
101	Lithuania	{lithuania}			Vilnius	Christianity	79.4	2866965	62680	66445	23844	0.875	Soviet Union	3.5	29.4	67.2	56.27	53.54	26.5	20.58	0	0	f	f	f
102	Niger	{niger}			Niamey	Islam	99.3	25369415	1266700	14915	591	0.4	France	41.6	19.5	38.7	23.31	11.42	16	0.07	0	0	f	t	f
103	Belarus	{belarus}			Minsk	Christianity	91	9200617	202900	68206	7121	0.808	Soviet Union	8.1	40.8	51.1	56.08	51.15	32.46	23.1	0	0	f	t	f
104	Georgia	{georgia}			Tbilisi	Christianity	88.1	3736400	69700	18696	4975	0.802	Soviet Union	8.2	23.7	67.9	43.35	41.03	46.4	41.3	0	0	f	f	f
105	Madagascar	{madagascar}			Antananarivo	Christianity	84.7	26923353	581540	14450	500	0.501	France	24	19.5	56.4	-11.57	-25.37	50.4	43.3	1	0	f	f	t
106	Paraguay	{paraguay}			Asunción	Christianity	96.1	6109644	397302	40458	6035	0.717	Spanish Empire	17.9	27.7	54.5	-19.17	-27.43	-54.15	-62.35	1	1	f	t	f
107	Venezuela	{venezuela}			Caracas	Christianity	92.6	28302000	882050	111813	3965	0.691	Spanish Empire	4.7	40.4	54.9	12.11	0.4	-59.45	-73.2	0	1	f	f	f
108	Malta	{malta}			Valletta	Christianity	88.5	519562	316	17721	33642	0.918	United Kingdom	1.1	10.2	88.7	36.05	35.47	14.34	14.12	0	0	f	f	t
109	Norway	{norway}			Oslo	Christianity	74.9	5514042	304282	482175	89242	0.961	Denmark	2.3	33.7	64	71.08	58	31.01	4.56	0	0	t	f	f
110	Thailand	{thailand}			Bangkok	Buddhism	90	68263022	510890	505982	7067	0.8		8.2	36.2	55.6	20.27	5.37	105.4	97.2	0	0	t	f	f
111	Afghanistan	{afghanistan}			Kabul	Islam	99.7	34262840	652867	14939	373	0.478	United Kingdom	23	21.1	55.9	38.22	29.23	75	60.3	0	0	f	t	f
112	Angola	{angola}			Luanda	Christianity	92.9	33086278	1246700	70533	2258	0.586	Portugal	10.2	61.4	28.4	-4.24	-18.02	24	11.4	1	0	f	f	f
113	Bulgaria	{bulgaria}			Sofia	Christianity	71.5	6447710	108612	84058	12207	0.795	Ottoman Empire	4.3	28	67.4	44.13	41.15	28.36	22.21	0	0	f	f	f
114	Republic of the Congo	{"republic of the congo",congo}			Brazzaville	Christianity	87.1	6106869	341500	0	2200	0.571	France	9.3	51	39.7	3.43	-5.02	18.4	11.1	2	0	f	f	f
115	The Gambia	{"the gambia",gambia}			Banjul	Islam	96.4	2417471	10000	0	772	0.5	United Kingdom	20.4	14.2	65.4	13.49	13.02	-13.48	-16.45	0	1	f	f	f
116	North Korea	{"north korea",dprk,"democratic people's republic of korea"}			Pyongyang	no religion	73	25660000	120538	16750	654	0	Empire of Japan	22.5	47.6	29.9	43	37.4	130.4	124.1	0	0	f	f	f
117	San Marino	{"san marino"}			San Marino	Christianity	0	33881	61	1702	50425	0		0.1	39.2	60.7	43.59	43.54	12.31	12.24	0	0	f	t	f
118	Tanzania	{tanzania}			Dodoma	Christianity	63.1	61741120	885800	70297	1136	0.549	United Kingdom	23.4	28.6	47.6	-0.59	-11.45	40.29	29.1	1	0	f	f	f
119	East Timor	{"east timor"}			Dili	Christianity	99.53	1354662	14874	2004	1517	0.607	Portugal	9.1	56.7	34.4	-8.19	-9.3	127.19	124	1	0	f	f	t
120	Equatorial Guinea	{"equatorial guinea"}			Malabo	Christianity	88.7	1558160	28051	12431	7605	0.596	Spain	2.5	54.6	42.9	3.47	-1.28	11.2	5.37	2	0	f	f	f
121	Eritrea	{eritrea}			Asmara		0	3748902	124330	2255	623	0.492	Ethiopia	11.7	29.6	58.7	18	12.22	43.7	36.5	0	0	f	f	f
122	Ivory Coast	{"ivory coast"}			Yamoussoukro		0	29389150	318003	69765	2539	0.55	France	20.1	26.6	53.3	10.44	4.22	-2.35	-8.35	0	1	f	f	f
123	Monaco	{monaco}			Monaco	Christianity	86	39050	2	8596	234317	0		0	14	86	43.45	43.44	7.26	7.24	0	0	t	f	f
124	Russia	{russia,"russian federation"}			Moscow	Christianity	61	146424729	16378410	1778782	12259	0.822		4.7	32.4	62.3	77.43	41.11	-169.4	19.38	0	0	f	f	f
125	Spain	{spain}			Madrid	non-practicing Catholic	37.5	48345223	498980	1427381	30058	0.905		2.6	23.2	74.2	43.47	36	3.19	-9.18	0	2	t	f	f
126	Switzerland	{switzerland}			Bern	Christianity	62.6	8902308	39997	812867	93525	0.962	Holy Roman Empire	0.7	25.6	73.7	47.48	45.5	10.29	5.57	0	0	f	t	f
127	Algeria	{algeria}			Algiers	Sunni Islam	99	45400000	2381741	163473	3700	0.745	France	13.3	39.3	47.4	37.05	18.58	12	-8.4	0	2	f	f	f
128	Azerbaijan	{azerbaijan}			Baku	Islam	0	10151517	86100	54622	5296	0.745	Soviet Union	6.1	53.5	40.4	41.53	38.24	50.33	44.46	0	0	f	t	f
129	Chile	{chile}			Santiago	Christianity	62.1	19960889	743812	317059	16265	0.855	Spanish Empire	4.2	32.8	63	-17.19	-53.53	-66.58	-75.38	1	1	f	f	f
130	Guatemala	{guatemala}			Guatemala City	Christianity	88	17602431	107159	85986	4883	0.627	Spanish Empire	13.3	23.4	63.2	17.49	13.44	-88.15	-92.15	0	1	f	f	f
131	New Zealand	{"new zealand"}			Wellington	no religion	48.6	5223100	262443	250451	48824	0.937		5.7	21.5	72.8	-34.23	-46.4	178.54	166.42	1	0	t	f	t
132	Turkmenistan	{turkmenistan}			Ashgabat	Islam	93	7057841	469930	53954	8508	0.745	Soviet Union	7.5	44.9	47.7	42.5	35.09	66.4	52.25	0	0	f	t	f
133	Belgium	{belgium}			Brussels	Christianity	63.7	11772639	30446	594104	51166	0.937	Netherlands	0.7	22.1	77.2	51.3	49.3	6.22	2.35	0	0	t	f	f
134	Cambodia	{cambodia}			Phnom Penh	Buddhism	97.1	17091464	176515	26669	1608	0.593	France	25.3	32.8	41.9	14.3	10.24	107.35	102.25	0	0	t	f	f
135	Haiti	{haiti}			Port-au-Prince	Christianity	87	11743017	27560	19044	1664	0.535	France	22.1	20.3	57.6	20.04	18.01	-71.38	-74.28	0	1	f	f	t
136	Netherlands	{netherlands}			Amsterdam	no religion	55.4	17960100	33893	1012847	57871	0.941	Spanish Empire	1.6	17.9	70.2	53.27	50.45	7.13	3.21	0	0	f	f	f
137	South Sudan	{"south sudan"}			Juba	Christianity	60.5	13249924	644329	4304	400	0.385	Sudan	0	0	0	12.14	3.29	35.19	23.23	0	0	f	t	f
138	Bahrain	{bahrain}			Manama	Islam	80.6	1577059	786	38869	26563	0.875	United Kingdom	0.3	39.3	60.4	26.17	25.34	50.43	50.23	0	0	t	f	t
139	Cyprus	{cyprus}			Nicosia	Christianity	72.3	918100	9241	28408	32281	0.896	United Kingdom	2	12.5	85.5	35.42	34.33	34.36	32.16	0	0	f	f	t
140	Nicaragua	{nicaragua}			Managua	Christianity	84.4	6733763	119990	14013	2046	0.667	Spanish Empire	15.5	24.4	60	15.01	10.43	-82.46	-87.4	0	1	f	f	f
141	Tunisia	{tunisia}			Tunis	Sunni Islam	99	11850232	155360	46687	3807	0.731	France	10.1	26.2	63.8	37.21	30.15	11.35	7.35	0	0	f	f	f
142	Germany	{germany}			Berlin	Christianity	0	84482267	348672	4259935	51073	0.942	West Germany	0.7	30.7	68.6	54.54	47.17	15.2	5.53	0	0	f	f	f
143	Guinea-Bissau	{guinea-bissau,"guinea bissau"}			Bissau	Islam	46.1	1781308	28120	1563	759	0.483	Portugal	50	13.1	36.9	12.39	10.55	-13.4	-16.42	0	1	f	f	f
144	Indonesia	{indonesia}			Jakarta	Islam	87	279118866	1811569	1186093	4333	0.705	Empire of Japan	13.7	41	45.4	6.04	-11	141.05	94.45	2	0	f	f	t
145	Lesotho	{lesotho}			Maseru	Christianity	92.1	2306000	30355	2373	1040	0.514	United Kingdom	5.8	39.2	54.9	-28.34	-30.4	29.28	27	1	0	t	t	f
146	Vatican City	{"vatican city",vatican}	UN observer			Christianity	0	764	0.49	0	0	0	Italy	0	0	0	41.54	41.54	12.27	12.26	0	0	t	t	f
147	Djibouti	{djibouti}			Djibouti		0	1001454	23180	3701	3348	0.509	France	2.4	17.3	80.2	12.41	10.56	43.25	41.46	0	0	f	f	f
148	Egypt	{egypt}			Cairo	Islam	0	105546000	995450	425906	3898	0.731	United Kingdom	11.7	34.3	54	31.36	22	36.15	24.4	0	0	f	f	f
149	Nauru	{nauru}			Yaren		0	11832	21	155	12390	0	United Kingdom	6.1	33	60.8	-0.3	-0.33	166.57	166.53	1	0	f	f	t
150	North Macedonia	{"north macedonia",macedonia}			Skopje	Christianity	60.4	1832696	25433	13881	6600	0.77	Yugoslavia	10.9	26.6	62.5	42.21	40.52	23	20.3	0	0	f	t	f
151	Ukraine	{ukraine}			Kyiv	Christianity	87.3	41130432	579300	200086	4596	0.773	Soviet Union	12.2	28.6	60	52.22	44.23	40.15	22.15	0	0	f	f	f
152	Honduras	{honduras}			Tegucigalpa	Christianity	75.7	9745149	111890	28489	2772	0.621	Spanish Empire	14.2	28.8	57	16.01	12.59	-83.35	-88.43	0	1	f	f	f
153	Syria	{syria}			Damascus	Islam	87	22923000	183630	19719	925	0.577	France	20	19.5	60.8	37.19	32.19	42.25	35.3	0	0	f	f	f
154	Togo	{togo}			Lomé	Christianity	47.8	8095498	54385	8160	944	0.539	France	28.8	21.8	49.8	11.08	6.07	1.46	-0.09	0	2	f	f	f
155	United Arab Emirates	{"united arab emirates",uae}			Abu Dhabi	Islam	76	9282410	83600	405468	43295	0.911	United Kingdom	0.9	49.8	49.2	26.05	22.38	56.24	51.33	0	0	t	f	f
156	Democratic Republic of the Congo	{"democratic republic of the congo"}			Kinshasa	Christianity	95.4	95370000	2267048	0	551	0.479	Belgium	19.7	43.6	36.7	5.22	-13.27	31.3	12.2	2	0	f	f	f
157	Liechtenstein	{liechtenstein}			Vaduz	Christianity	83.2	39680	160	6608	169260	0.935	German Confederation	7	41	52	47.14	47.03	9.33	9.28	0	0	t	t	f
158	Luxembourg	{luxembourg}			Luxembourg City	Christianity	73.2	660809	2586	85506	133745	0.93		0.3	12.8	86.9	50.11	49.27	6.32	5.44	0	0	t	t	f
159	Mauritius	{mauritius}			Port Louis	Hinduism	48.54	1261041	2030	11525	8873	0.803	United Kingdom	4	21.8	74.1	-10.25	-20.32	63.3	56.35	1	0	f	f	t
160	Morocco	{morocco}			Rabat	Islam	99.6	37150700	446300	142867	3853	0.683	France	14	29.5	56.5	35.55	27.4	-1.1	-13.15	0	1	t	f	f
161	Myanmar	{myanmar}			Naypyidaw	Buddhism	87.9	55770232	653508	58582	1089	0.585	United Kingdom	24.1	35.6	40.3	28.32	9.59	101.1	92.15	0	0	f	f	f
162	Nigeria	{nigeria}			Abuja		0	216783400	910768	430923	2019	0.535	United Kingdom	21.1	22.5	56.4	13.5	4.17	13.15	2.45	0	0	f	f	f
163	Slovenia	{slovenia}			Ljubljana	Christianity	77.8	2117674	20151	61749	29135	0.918	Yugoslavia	1.8	32.2	65.9	46.52	45.25	16.36	13.23	0	0	f	f	f
164	Bolivia	{bolivia}			Sucre	Christianity	89.3	12006031	1083301	40408	3345	0.692	Spain	13.8	37.8	48.2	-9.4	-22.54	-57.27	-69.4	1	1	f	t	f
165	India	{india}			New Delhi	Hinduism	79.8	1392329000	2973190	3201471	2274	0.633	United Kingdom	15.4	23	61.6	37.06	8.06	97.23	68.03	0	0	f	f	f
166	Libya	{libya}			Tripoli	Islam	0	6931061	1759540	39006	5791	0.718	United Kingdom	1.3	52.3	46.4	33.1	19.3	25.05	9.2	0	0	f	f	f
167	Maldives	{maldives}			Malé	Islam	0	382751	298	5406	10366	0.747	United Kingdom	3	16	81	7.07	-0.42	73.31	72.42	2	0	f	f	t
168	Mongolia	{mongolia}			Ulaanbaatar	Buddhism	51.7	3457548	1553556	15098	4510	0.739	Qing China	12.1	38.2	49.7	52.09	41.34	119.56	88	0	0	f	t	f
169	Suriname	{suriname}			Paramaribo	Christianity	52.3	616500	156000	3224	5259	0.731	Netherlands	11.6	31.1	57.4	6.01	1.5	-54	-58.05	0	1	f	f	f
170	Sweden	{sweden}			Stockholm	Christianity	61.4	10548336	407284	635664	60730	0.947	Kalmar Union	1.6	33	65.4	69.04	55.2	24.15	10.57	0	0	t	f	f
171	United Kingdom	{"united kingdom",uk,"great britain",gb}			London	Christianity	0	67026292	241930	3131378	46542	0.929		0.7	20.2	79.2	58.4	49.57	1.46	-7.34	0	2	t	f	t
172	Ethiopia	{ethiopia}			Addis Ababa	Christianity	67.3	107334000	1096630	99269	825	0.498		34.8	21.6	43.6	14.53	3.24	48	33	0	0	f	t	f
173	Italy	{italy}			Rome	Christianity	84.4	58780965	294140	2107703	35579	0.895		2.1	23.9	73.9	47.05	37.56	18.31	6.37	0	0	f	f	f
174	Kiribati	{kiribati}			South Tarawa	Christianity	96.2	120740	811	227	1765	0.624	United Kingdom	23	7	70	3.23	-11.26	-150.13	169.32	2	0	f	f	t
175	Latvia	{latvia}			Riga	Christianity	64	1882000	62249	39854	21267	0.863	Soviet Union	3.9	22.4	73.7	58.05	55.4	28.15	21	0	0	f	f	f
176	Palau	{palau}			Ngerulmud	Christianity	78.7	16733	459	218	12084	0.767	United States	3	19	78	8.1	2.47	134.42	131.07	0	0	f	f	t
177	South Africa	{"south africa",rsa}			style	Christianity	78	62027503	1214470	419016	7055	0.713		2.8	29.7	67.5	-22.2	-34.5	32.53	16.3	1	0	f	f	f
178	Australia	{australia}			Canberra	Christianity	43.9	26853200	7633565	1734532	66916	0.951		3.6	25.3	71.2	-10.41	-39.08	153.38	113.09	1	0	t	f	f
179	Estonia	{estonia}			Tallinn	no religion	58.4	1365884	42388	37191	27991	0.89	Soviet Union	2.8	29.2	68.1	59.4	57.31	28.12	23.24	0	0	f	f	f
180	Japan	{japan}			Tokyo		0	124340000	364546	4940878	39650	0.925		1.1	30.1	68.7	45.31	30.59	145.49	129.43	0	0	t	f	t
181	Rwanda	{rwanda}			Kigali	Christianity	93.8	13246394	24668	11070	822	0.534	Belgium	30.9	17.6	51.5	-1.03	-2.51	30.53	28.52	1	0	f	t	f
182	Sierra Leone	{"sierra leone"}			Freetown	Islam	0	8494260	71620	4249	505	0.477	United Kingdom	60.7	6.5	32.9	9.59	6.56	-10.17	-13.18	0	1	f	f	f
\.


--
-- Data for Name: ethnic_groups; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.ethnic_groups (country_id, name, percentage) FROM stdin;
67	Khas Arya	31.25
70	Mestizo	52.9
73	Bambara	33.3
74	Punjabis	38.78
78	 Andorrans	48.3
86	Montenegrins	45
90	Mestizo	60.2
97	White	61.6
102	Hausa	53.1
107	Moreno	51.6
107	White	43.6
110	Thai (Central Thai)	37
111	Pashtun	42
112	Ovimbundu	37
115	Mandinka	34.4
121	Tigrinya	50
122	Akan	38
130	Ladino	56.01
130	Maya	41.66
133	Belgians	66.6
138	Arabs	53.2
138	Asians	43.4
140	Mestizo	69
150	Macedonians	58.4
155	South Asian	59.4
155	Indian	38.2
157	Liechtensteiners	66.2
159	Indian	67
160	Arabs	67
160	Berbers	31
164	Mestizo	68
161	Bamar	68
172	Oromo	34.5
175	Latvians	63
176	Palauan	65.2
176	Asian	31.6
182	Mende	35.5
182	Temne	33.2
4	Visayan	33.7
10	Persians	0
10	Azeris	0
10	Kurds	0
10	Mazanderanis	0
10	Lurs	0
10	Gilaks	0
10	Arabs	0
10	Armenians	0
10	Turkmens	0
11	Mestizo	65
14	Haratin	40
17	Bosniaks	50.1
17	Serbs	30.8
18	White	47.7
18	Mixed	43.1
19	Mossi	52
22	White	64.1
23	Fula	33.4
30	Akans	45.7
35	Fon	38.4
38	Chuukese	48.8
42	Wolof	39.7
50	Ovambo	49
54	Indian	39.8
55	Lao	53.2
49	Bumiputera	69.7
49	Malay	57.3
65	Kuwaiti	32.59
65	Asian	40.3
59	Arab	40
59	South Asian	36
59	South Asian	36
\.


--
-- Data for Name: funfacts; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.funfacts (country_id, text) FROM stdin;
55	This country is very nice and cute and communist like an anime girl with a flag of Soviet Union on her wall
60	This country is one of the most fragile states (failed states)
75	This country experienced hyperinflation with the most extreme inflation rate in history
75	This country has been under the communist rule
136	Some people think you would not get the whole experience of being in this country if you stay sober
136	This country doesn't want to be associated with drugs
136	For some tourists, this country's capital is kinda sinonym to the whole country
112	This country has been under the communist rule
31	They are so mad about linguistic purism in this country that you would face problems with getting a passport for your child if you don't name them with one of the registered national names
31	This country's official language is so well saved that people who speak it can read papers that are 500 years old without any knowledge of ancient languages
76	Most people think you would not get the whole experience of being in this country if you stay sober
164	One of the most famous revolutionists died in this country
134	This country was attempted to be turned in communist paradise (literally), but everything went wrong, because there is no communist paradise possible
134	This country has been under the communist rule
134	There is a niche music genre originated from this country which became known because of mass terror that happened there in past
1	This country is one of the most fragile states (failed states)
1	This country is like a responsible religious parent whose crazy rebellious child has run away from home
129	One of geographic regions of this country is known to be the dryest place in the world
129	According to the legend, the first man who did a trip around the world named Pacific Ocean Pacific because of difficulties he experienced while being in waters of this country
47	Ideally, you should travel to the certain place in this country at least once to be good in its official religion
48	This country was the first country to fully legalise recreational use of marijuana
20	The whole youth movement (usually assosiated with New Age and LSD) exists because of this country
20	This country has been under the communist rule
66	In this country, a famous humanitarian and Nobel Peace Prize laureate was born. This person founded an international religious congregation helping people in need (poor, sick, orphaned children, etc.)
66	This country has been under the communist rule
142	This country used to be very powerful and very evil. Now it's still powerfull, but nice and fluffy.
142	Once I seriously considered applying to the Faculty of Degustation in the University of Beer. And I knew where I needed to go for this.
142	Despite the common stereotype, trains in this country are not punctual at all
142	This country has been under the communist rule
14	When I was a kid, I liked eating with my hands, and when my mother would say it's not good, I would answer: "In <this country>, people eat everything with their hands!". Although this won't give you much information since there are lot of countries where people eat with their hands, it's still a fact that in this country people eat with their hands.
57	This country is one of the most fragile states (failed states)
57	This country can be an ideal for all aspiring anarcho-capitalists
57	This country has been under the communist rule
40	This country is a birthplace of a person that is believed to be a prophet in some religions
40	Knowing some botany is crucial when visiting this country
40	Some old people are starting to sing when they hear the name of this country
41	People drink A LOT in this country (and it's clear why, they produce great vine)
41	This country has been under the communist rule
107	If you want to buy some groceries with cash in this country, you will have to bring a suitcase full of money like you are a drugdealer
104	What??? Is it a country? I thought it's just one of the States...
104	They have really good vine there
104	This country has been under the communist rule
171	This country had a huge influence on almost all kinds of rock music
171	This country can be treated as one of the countries which had the most historical significance to the world as we know it now
171	People of this country may still be treated as arrogant and superior because of this country's history
97	This country can be treated as one of the countries which had the most historical significance to the world as we know it now
97	There are a lot of weird religious movements in this country
97	This country is extremely diverse in terms of everything: cultures, languages, ethnicities, serial killers, etc.
97	Fentanyl, crack, crystal meth. This country is known for using these things a lot.
97	There was one programmer in this country that created his own operating system because God told him to do so
113	This country has been under the communist rule
113	Yoghurt was invented in this country
85	Almost no one is interested in this country now, but its former glory fascinates a lot of people around the world
85	A lot of terms we use now are coming from an ancestor of this country's language
54	This country was a place for a well-known totalitarian sect where people collectively commited suicide
161	This country is one of the most fragile states (failed states)
161	The colonial name of this country starts with B
26	This country is the most linguistically dense place in the world
26	They eat human flesh there. Even ChatGPT says it's true.
90	If you want to connect to spirits, talk to you ancestors and get lost in your ego in an authentic way, you should visit this country
38	On the island that now belongs to this country, large rock discs with holes in the middle of them were used as money by its inhabitants.The ownership of these rocks was recorded in oral history. Some people claim that this system was a prototype of blockchain.
38	This country depends hardly on the economic help from the US
124	The most stereotypical image of this country is a wild animal drinking an alcoholic drink and playing a musical instrument
124	This country's population used a lot of desomorphine in the beginning of the 21st century
124	It seems this country's government is paranoid, which is kinda dangerous for the world
124	This country has been under the communist rule
169	This country has the most percentage of forest area in the world
98	This country is one of the most fragile states (failed states)
98	When you hear the name of this country, you might think of a strong, based, very masculine alfa male
98	There is a large lake named like this country
156	This country is one of the most fragile states (failed states)
156	If you are radical at fighting against child labour, you should throw away your phone, because raw materials used for it were mined by children of this country in terrible working conditions
22	If you are a woman in this country, you are probably dreaming of getting married to a tourist
22	National hero of this country turned to a successfull brand now, which is very ironic
52	One of the most famous revolutionists was born in this country
52	This country experiences an inflation crisis now
52	They love meat a lot in this country
149	This country severely suffered from phosphate mining, which damaged its ecosystems and even made its area decrease
149	This country is a client state of Australia
177	In this country, ordinary people usually have metal doors, a lot of locks on them, signalisation and other secutity measures for their houses
177	This country more than one capital
53	This country is known for its cat women and a great working nation
53	It is really a problem to use VPN here
53	This country is very haram in terms of livestock farming
39	A piece of plant often accosiated with this country and depicted on its coat of arms looks like a pussy
39	There are a lake and a waterfall named like this country's capital
146	One of former official languages of this country is officially extinct
146	You cannot be a citizen of this country without being a citizen of an other country
103	POTATOS!!!
72	One good old system programmer was born in this country
72	There is a national sweet in this country that one either hates or loves
24	Security is the biggest concern in this country, so everybody should serve in the military there
24	The military conflict between this country and its neigbouring territories... you don't have to read further, just type the country
110	I can't resist making jokes about some unusual girls here
151	If you say anything bad about this country, you would be cancelled in a nanosecond
151	Flag of this country illustrates its typical landscape with reference to its black soil wealth
151	This country is home to the creators of some snuff videos that were really popular in 2007
151	This country has been under the communist rule
180	One of stereotypical images of this country is a single man living with his parents, never leaving his room, playing computer games all the time and thinking about suicide
180	The term meaning "little girl" that is going from this country's language is originally a name of an American girl from a novel by a Russian writer
111	One of promoters of psychedelics was caught in this country's capital and sent to prison in the US
111	There is a huge debate in the world whether women feel happy in this country
111	This country has been under the communist rule
111	This country is one of the most fragile states (failed states)
111	This country is number one in the world by opiates use
16	This country was the first country visited by Colombus in his first voyage
181	This country is known for a mass genocide happening there in the 20th century
102	If you pronounce the name of this country incorrectly, you would be cancelled in a nanosecond
102	One letter divides the name of this country from the word you could be fired for
109	Some people here are really fond of burning churches
153	This country is one of the most fragile states (failed states)
153	This country is a dream destination for all young and strong Jihad lions, even though the Khalifah now lost its power
95	Due to the recent concerns about sea level rise this country announced plans to move all its economy and culture to the metaverse
32	The former leader of this country was "Lord of All the Beasts of the Earth and Fishes of the Seas". Ah, what a great ruler we lost!
13	This country has been under the communist rule
13	This country's capital is very old and historically rich city
173	They sell a lot of expensive clothes
173	This country's national food is a subject of millions of jokes that are even not funny anymore
15	Capital of this country has been under the rule of a social dempcratic party for more than a decade, which influenced architecture there
93	This country was the first to legalize shrooms
93	This country used to be a powerful empire, but now it's kind of a shithole compared to other countries in the region
174	Eastern parts of this country have the earliest time on Earth (UTC+14), so this country should be called "Land of the Rising Sun"
174	If you want to disappear without a trace, just visit this country. You could hardly find more of a shithole.
101	This country's language is very old and well-saved compared to other languages of the region
101	This country has been under the communist rule
35	This country has been under the communist rule
168	This country used to be a huge empire, but now some people even don't know it still exists
168	This country has been under the communist rule
28	This country has been under the communist rule
28	This country is often confused with a similarly named country, and actually no one cares, both countries are not that significant in the world economy to care about them
170	Some people here are really fond of burning Quran
170	You sometimes might mention this country while talking about sexual education
170	This country gave us standards of how to make good pop music and good cheap furniture
165	People in this country are so tolerant to noise and crowdedness that they can sleep everywhere
165	Mythology and religion of this country has been taken as a reference infinite number of times creating infinite number of cults, sects, and some random esoteric movements
165	Symbol for which you can be fired in the West is used very widely in this country
165	Speakers of the most Western Austroasian languages live in this country
125	There is an interesting nation living in this country. Their language has no relatives at all.
125	These guys love watching at animals being killed in front of the crowd. Crazy people.
125	This country used to be a very strong and evil empire, but now they are nice and fluffy
67	One of the world's most famous religious teachers was born here
67	This country is probably the best place for mountain hiking
67	This country has a really interesting and unique flag
116	It seems this country's government is crazy, which might be kinda dangerous for the world
116	Businesswomen are more common in this country than businessmen due to some law details
116	If you know what is Kwangmyong then you guessed the country
116	One can hardly think of anything positive about this country
126	One good old chemist once synthesised a substance in his laboratory in this country, and the world has gone mad...
126	When I was a child I used to think that you need to be a king or something to live in this country. Maybe I still think so.
58	One of banknotes printed in this country during the period of hyperinflation is cosidered a banknote with the biggest number of zeroes ever in history
114	This country has been under the communist rule
114	This country is named after a river. A large, deep, and long river.
10	ISIS hates people in this country even thougn they are muslims
10	They recently gave a Noble Peace Prize to a political activist from this county
25	There is a weapon on the flag of this country, and it looks very cool, you can easily imagine a child running with this weapon through the jungle
25	This country has been under the communist rule
5	A big war conflict started with invasion in this country
5	This country's language is very funny to internet users
100	This country has a really weird power-sharing arrangement based on religions. Each religious sect is assigned a specific number of seats in the country' s parliament.
137	This country is one of the most fragile states (failed states)
137	This country can be an ideal for all aspiring anarcho-capitalists
137	This country is like a rebelious 14-year-old child running away from his parents
172	A representation of God on Earth was born in this country, and now people from another side of the world are travelling there to return to their roots
172	This country is believed to be the origin of one top predator spieces that now has spread all over the world
172	This country has been under the communist rule
172	This country has never been colonised (although there were some silly attempts)
175	They really hate Soviet Union in this country
175	This country has been under the communist rule
132	The president of this country prohibited all cars that are not white. Also he wants his portait to be hung in every room, even in front of a toilet. Classic never gets old.
132	This country has been under the communist rule
135	In this country people practice vodoo religion and eat human flesh!!! (though ChatGPT thinks there is no cannibalism). Aaaa!!!
\.


--
-- Data for Name: languages; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.languages (country_id, name) FROM stdin;
57	Somali
79	Bengali
168	Mongolian
25	Portuguese
153	Arabic
44	English
44	French
44	Bislama
66	Albanian
70	English
67	Nepali
26	English
26	Hiri Motu
26	PNG Sign Language
26	Tok Pisin
60	Arabic
128	Azerbaijani
113	Bulgarian
134	Khmer
18	Portuguese
65	Standard Arabic
96	Arabic
170	Swedish
81	Spanish
82	French
10	Persian
120	French
120	Portuguese
120	Spanish
40	English
49	Malay
38	English
111	Persian (Dari)
111	Pashto
112	Portuguese
35	French
118	Swahili
118	English
132	Turkmen
34	Uzbek
160	Arabic
160	Berber
136	Dutch
109	Sami languages
178	English
14	Arabic
152	Spanish
24	Hebrew
158	French
158	German
3	English
3	Chichewa
50	English
88	English
88	French
12	Portuguese
69	Sinhala
69	Tamil
61	Armenian
53	Standard Chinese
125	Spanish
106	Spanish
106	Guaraní
138	Arabic
103	Russian
167	Dhivehi
131	English
131	Maori language
164	Castilian
164	Aymara
164	Araona
164	Baure
164	Bésiro
164	Canichana
164	Cavineña
164	Cayubaba
164	Chácobo
164	Chimán
164	Ese Ejja
164	Guaraní
164	Guarasu'we
164	Guarayu
164	Itonama
164	Leco
164	Machajuyai-Kallawaya
164	Machineri
164	Maropa
164	Mojeño-Ignaciano
164	Mojeño-Trinitario
164	Moré
164	Mosetén
164	Movima
164	Pacawara
164	Puquina
164	Quechua
164	Sirionó
164	Tacana
164	Tapieté
164	Toromona
164	Uru-Chipaya
164	Weenhayek
164	Yaminawa
164	Yuki
164	Yuracaré
164	Zamuco
54	English
37	Russian
39	English
39	French
126	German
126	French
126	Italian
126	Romansh
62	French
62	Sango
143	Portuguese
162	English
51	English
51	Malay
51	Mandarin Chinese
51	Tamil
129	Chilean Spanish
63	Spanish
122	French
174	English
174	Gilbertese
175	Latvian
56	English
20	Vietnamese
127	Arabic
127	Berber
148	Arabic
30	English
89	Spanish
89	Quechua
144	Indonesian
180	Japanese
15	German
133	Dutch
133	French
133	German
147	Arabic
147	French
116	Korean
176	English
181	English
181	French
181	Kinyarwanda
181	Swahili
21	Serbian
121	Tigrinya
142	German
73	Bambara
73	Bobo
73	Bozo
73	Dogon
73	Fula
73	Hassaniya
73	Kassonke
73	Maninke
73	Minyanka
73	Senufo
73	Songhay languages
73	Soninke
73	Tamasheq
47	Arabic
94	Tajik
43	Turkish
32	English
32	Swahili
155	Arabic
156	French
123	French
102	French
36	French
77	Korean
77	Korean Sign Language
171	English
4	Filipino
4	English
42	French
52	Spanish
45	Croatian
172	Afar
172	Amharic
172	Oromo
172	Somali
172	Tigrinya
74	Urdu
74	English
78	Catalan
139	Greek
139	Turkish
99	Arabic
99	Kurdish
23	French
135	French
7	English
7	Swahili
6	Romanian
182	English
83	Dzongkha
98	Arabic
98	French
22	Spanish
108	English
161	Burmese
90	Quechua
90	Aymara
90	Asháninka
124	Russian
177	Afrikaans
177	English
177	Sotho
177	Northern Sotho
177	Swazi
177	Tsonga
177	Tswana
177	Venda
177	Xhosa
177	Zulu
91	English
91	French
9	English
9	Swazi
166	Arabic
105	French
105	Malagasy
163	Slovene
1	Arabic
1	English
141	Arabic
95	English
80	Spanish
179	Estonian
173	Italian
71	English
130	Spanish
165	Hindi
165	English
154	French
58	Chewa
58	Chibarwe
58	English
58	Kalanga
58	Khoisan
58	Nambya
58	Ndau
58	Ndebele
58	Shangani
58	Shona
58	Sign language
58	Sotho
58	Tonga
58	Tswana
58	Venda
58	Xhosa
75	Hungarian
140	Spanish
151	Ukrainian
107	Spanish
33	English
46	Danish
119	Portuguese
119	Tetum
85	Greek
100	Arabic
28	Slovak
27	Malay
64	Russian
55	Lao
93	Portuguese
117	Italian
97	English
92	Arabic
157	German
150	Macedonian
150	Albanian
87	English
48	None (Spanish has de facto status)
146	Italian
146	Latin (formerly)
5	Polish
101	Lithuanian
41	Romanian
149	English
110	Thai
84	Spanish
2	Arabic
2	French
145	Sotho
145	English
13	Slovak
72	Swedish
159	Mauritian Creole
8	None (Spanish has de facto status)
11	Spanish
19	French
29	French
29	Kirundi
29	English
114	French
59	Arabic
137	English
169	Dutch
\.


--
-- Data for Name: schema_migrations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.schema_migrations (version, dirty) FROM stdin;
4	f
\.


--
-- Name: countries countries_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.countries
    ADD CONSTRAINT countries_pkey PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: ethnic_groups ethnic_groups_country_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ethnic_groups
    ADD CONSTRAINT ethnic_groups_country_id_fkey FOREIGN KEY (country_id) REFERENCES public.countries(id);


--
-- Name: funfacts funfacts_country_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.funfacts
    ADD CONSTRAINT funfacts_country_id_fkey FOREIGN KEY (country_id) REFERENCES public.countries(id);


--
-- Name: languages languages_country_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.languages
    ADD CONSTRAINT languages_country_id_fkey FOREIGN KEY (country_id) REFERENCES public.countries(id);


--
-- PostgreSQL database dump complete
--


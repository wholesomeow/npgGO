COPY ngram_fantasy (
    NGram_ID,
    NGram_VAL,
    NGram_POS
) FROM 'npgGO/database/rawdata/csv/Fantasy_Names_NGrams.csv' DELIMITER ',' CSV HEADER;
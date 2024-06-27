-- name: InsertReview :one
INSERT INTO public.reviews
(product_id, customer_id, rating, review_text, created_at, last_modified_at, deleted_at)
VALUES($1, $2, $3, $4, $5, $6, $7)
ON CONFLICT ON CONSTRAINT reviews_product_customer
DO UPDATE SET
    review_text = EXCLUDED.review_text,
    rating = EXCLUDED.rating,
    last_modified_at = EXCLUDED.last_modified_at,
    deleted_at = EXCLUDED.deleted_at
RETURNING id;

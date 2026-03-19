#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_empty() {
        let data = vec![];
        assert_eq!(calculate_sum_of_squares(&data), 0);
    }

    #[test]
    fn test_positive_numbers() {
        let data = vec![1, 2, 3];
        assert_eq!(calculate_sum_of_squares(&data), 14);
    }

    #[test]
    fn test_negative_numbers() {
        let data = vec![-1, -2, -3];
        assert_eq!(calculate_sum_of_squares(&data), 14);
    }
}
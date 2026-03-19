#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_new_counter() {
        let counter = Counter::new(10);
        assert_eq!(counter.get(), 10);
    }

    #[test]
    fn test_increment() {
        let mut counter = Counter::new(0);
        counter.increment();
        assert_eq!(counter.get(), 1);
    }

    #[test]
    fn test_multiple_increment() {
        let mut counter = Counter::new(5);

        for _ in 0..3 {
            counter.increment();
        }

        assert_eq!(counter.get(), 8);
    }
}
use rand::{distributions::Uniform, Rng}; // 0.6.5

// Generates practice vector
// elements variable indicates number of elements in the vector
pub fn random_vector
    <T: rand::distributions::uniform::SampleUniform>
    (lowest: T, highest: T, elements: i64)
    -> Vec<T>
{
    let mut rng = rand::thread_rng();
    let range = Uniform::new(lowest, highest);

    (0..elements).map(|_| rng.sample(&range)).collect()
}

pub fn print_type_of<T>(_: T) {
    println!("{}", std::any::type_name::<T>())
}
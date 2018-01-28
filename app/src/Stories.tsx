import * as React from 'react';
import { connect, Dispatch } from 'react-redux';
import { RootState } from './store';
import { Story } from './reducers/stories';
import { Container, Grid, Header } from 'semantic-ui-react';
import StoryList from './StoryList';
import StoryView from './StoryView';
import { RootAction } from './actions';
import { listStories } from './actions/stories';

type StoriesProps = {
  stories: Story[],
  loading: boolean,
  error: Error | null,
  dispatch: Dispatch<RootAction>
};

class Stories extends React.Component<StoriesProps, {}> {

  componentDidMount() {
    this.props.dispatch(listStories());
  }

  render() {
    console.log(this.props.stories)
    return (
      <Container style={{padding: '1em'}} fluid={true}>
        <Header as="h1" dividing={true}>Hacker News with gRPC-Web</Header>

        <Grid columns={2} stackable={true} divided={'vertically'}>
          <Grid.Column width={4}>
            <StoryList stories={this.props.stories}/>
          </Grid.Column>

          <Grid.Column width={12} stretched={true}>

            <Header as="h2">Example body text</Header>
            <StoryView/>
          </Grid.Column>
        </Grid>

      </Container>
    );
  }

}

export default connect((state: RootState) => ({
  stories: Object.keys(state.stories.stories).map(key => state.stories.stories[key]),
  loading: state.stories.loading,
  error: state.stories.error,
}))(Stories);
